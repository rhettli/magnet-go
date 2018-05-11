package controller

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"fmt"
	"os/exec"
	"bytes"
	"net/http"
	"encoding/json"
	"magnet/model"
	"magnet/mdb"
	"net/url"
	"path/filepath"
	"os"
	"log"
	"strings"
	"runtime"
)

const phpPathWindows="f:/phpStudyo/php/php-5.4.45/php"
const phpPathLinux ="php"
//const phpExecFile  = `C:\Users\Oshine\Desktop\cloudbooks.top\controller\demo.php`

func GetPhpCgiBySystem()string{
 	if runtime.GOOS=="windows"{
		return phpPathWindows
	}
	return phpPathLinux
}


func GetDetailApi(id string) (*model.MiltTbHashDetai ,error) {

	db,err:=mdb.GetLocalDB()
	if err!=nil{
		fmt.Println(err)

		return nil,err
	}

	defer db.Close()

	sqlField:="file_list,search_hash.info_hash,search_hash.name,search_hash.create_time,search_hash.last_seen,search_hash.length,search_hash.requests "
	sqlJoin:="join search_filelist on  search_hash.id ="+id+" and search_filelist.info_hash=search_hash.info_hash"

	miltHash:=model.MiltTbHashDetai{}

	db.Table("search_hash").Select(sqlField).Joins(sqlJoin).Scan(&miltHash)

	//fmt.Println("--",miltHash)

	//jdata,err:=json.Marshal(miltHash)

	return &miltHash,nil
}



func SearchHashApi(c *gin.Context) {

	key:=c.Params.ByName("key")
	page:=c.Params.ByName("page")

	if _,err:=strconv.Atoi(page);err!= nil {//不是数字格式的，则默认为第一页
		page="1"
	}

	//url编码
	resUri, pErr := url.Parse(key)
	if pErr!=nil{
		c.JSON(200,nil)
		return
	}
	fmt.Println("searchHashApi",resUri.EscapedPath(),page)
	//函数返回一个*Cmd，用于使用给出的参数执行name指定的程序

	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	path:=strings.Replace(dir, "\\", "/", -1)

	path=path+"/phpext/sphinxsearch.php"
	fmt.Println(path)

	cmd := exec.Command(GetPhpCgiBySystem(),path,resUri.EscapedPath(),page)

	//读取io.Writer类型的cmd.Stdout，再通过bytes.Buffer(缓冲byte类型的缓冲器)将byte类型转化为string类型(out.String():这是bytes类型提供的接口)
	var out bytes.Buffer
	cmd.Stdout = &out

	//Run执行c包含的命令，并阻塞直到完成。  这里stdout被取出，cmd.Wait()无法正确获取stdin,stdout,stderr，则阻塞在那了
	err = cmd.Run()
	if err!= nil{
		fmt.Println(err.Error()+out.String())
		c.JSON(400,err.Error()+out.String())
		return
	}

	//fmt.Println("backjosn:",out.String())

	var part model.SPhinxHash
	if err := json.Unmarshal([]byte(out.String()), &part); err != nil {
		fmt.Println("string to josn error:",err)
		fmt.Println("josntring:",out.String())
		c.JSON(http.StatusOK,nil)
		return
	}else {
		fmt.Println("================json str 转struct==")

		db,err:=mdb.GetLocalDB()
		if err!=nil{
			fmt.Println(err)
			c.JSON(http.StatusOK,nil)
			return
		}
		defer db.Close()

		for ind,item:=range part.Matches{//远程数据没有name 和 requests，循环获取每一个匹配项目并获取name和requests

			findSearchHash:=model.SearchHash{}

			if err=db.Where(&model.SearchHash{Id: item.Attrs.Hash_id.String()}).First(&findSearchHash).Error ;err!=nil{
				fmt.Println(err)
				c.JSON(http.StatusOK,nil)
				return
			}
			//注意目前http 的header中加入中文的话会乱码，这里urlencode编码下
			resName, pErr := url.Parse(findSearchHash.Name)
			if pErr!=nil{
				c.JSON(200,nil)
				return
			}
			//fmt.Println("searchHashApi",resUri.EscapedPath(),page)

			part.Matches[ind].Attrs.Name=resName.EscapedPath() 		//findSearchHash.Name
			part.Matches[ind].Attrs.Requests=findSearchHash.Requests
		}

		//json数据转字符串
		jStr,err:=json.Marshal(part)
		if err!=nil { //json转字符串错误
			fmt.Println(err)
			c.JSON(http.StatusOK, nil)
			return
		}

		//c.Header("Content-disposition","attachment;filename=中国.txt")
		//把结果放到header中的backstr中去
		c.Header("backstr",string(jStr))
		c.JSON(http.StatusOK,nil)
		fmt.Println("================json str 转struct over==")
	}


}

