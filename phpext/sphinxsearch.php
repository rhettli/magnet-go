<?php

include "sphinxapi.php";


const PAGE_COUNT=15;
const MODE = SPH_MATCH_ALL;
const RANKER = SPH_RANK_WORDCOUNT;
const HOST = "118.25.10.97";
const PORT = 9312;
const INDEX = "*";


$cl = new SphinxClient ();

$cl->SetServer(HOST, PORT);
$cl->SetConnectTimeout ( 1 );
$cl->SetArrayResult ( true );
$cl->SetWeights ( array ( 100, 1 ) );
$cl->SetMatchMode ( MODE );


//if ( count($filtervals) )	$cl->SetFilter ( $filter, $filtervals );
//if ( $groupby )				$cl->SetGroupBy ( $groupby, SPH_GROUPBY_ATTR, $groupsort );
//if ( $sortby ){
//    $cl->SetSortMode ( SPH_SORT_EXTENDED, $sortby.' desc' );
    //if ( $sortexpr )
    //$cl->SetSortMode ( SPH_SORT_EXPR,'desc');
//}
//if ( $distinct )			$cl->SetGroupDistinct ( $distinct );
//if ( $select )				$cl->SetSelect ( $select );

$page=isset($argv[2])?$argv[2]:"1" ;

$cl->SetLimits(((int)$page-1)*PAGE_COUNT,PAGE_COUNT,1000); //( 0, $limit, ( $limit>1000 ) ? $limit : 1000 );

$cl->SetRankingMode ( RANKER );

$key=$argv[1];

$res = $cl->Query (urldecode($key), INDEX );

$res["kv"]="--".urldecode( $key )."-".$argv[2];

if ($res===false ){
    print "Query failed: " . $cl->GetLastError() . ".\n";
} else {
    //if ( $cl->GetLastWarning() )	print "WARNING: " . $cl->GetLastWarning() . "\n\n";
    //获取所有的页码
    //$this->total_found= $res['total_found'];

    $data['total_found']=$res['total_found'];

    //print $res['total_found'];

    print (json_encode( $res));
}

