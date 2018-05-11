// forked from http://code.9leap.net/codes/show/4497

'use strict';

enchant();
var SIZE = 1070;
var TILE_SIZE = SIZE / 4;

window.onload = function () {
  var game = new Game(SIZE, SIZE);
  game.fps = 24;
  game.preload('http://www.cloudbooks.top/static/img/a.jpg');

  // 乱数発生関数
  var rand = function rand(num) {
    return Math.random() * num | 0;
  };

  // グリッドライン用グループ
  var Grid = enchant.Class.create(Group, {
    initialize: function initialize() {
      Group.call(this);
      for (var i = 0; i < 4; i++) {
        var leftLine = new Line(i % 4 * TILE_SIZE, 0);
        this.addChild(leftLine);
        var rightLine = new Line(i % 4 * TILE_SIZE + TILE_SIZE - 1, 0);
        this.addChild(rightLine);
        var topLine = new Line(TILE_SIZE * 2, i % 4 * TILE_SIZE - TILE_SIZE * 2);
        topLine.rotate(90);
        this.addChild(topLine);
        var bottomLine = new Line(TILE_SIZE * 2, i % 4 * TILE_SIZE - TILE_SIZE - 1);
        bottomLine.rotate(90);
        this.addChild(bottomLine);
      }
    }
  });

  //グリッドライン用ライン
  var Line = enchant.Class.create(Sprite, {
    initialize: function initialize(x, y) {
      Sprite.call(this, 1, SIZE);
      var surface = new Surface(1, SIZE);
      var ctx = surface.context;
      ctx.strokeStyle = '#000000';
      ctx.beginPath();
      ctx.moveTo(0, 0);
      ctx.lineTo(0, SIZE);
      ctx.stroke();
      this.image = surface;
      this.x = x;
      this.y = y;
    }
  });

  //パネル
  var Panel = enchant.Class.create(Sprite, {
    initialize: function initialize(width, height) {
      Sprite.call(this, width, height);
      this.image = game.assets['http://www.cloudbooks.top/static/img/a.jpg'];
      this.position = 0;
      this.moved = false;
      this.addEventListener(enchant.Event.TOUCH_START, this.onTouchStart);
      this.addEventListener(enchant.Event.TOUCH_END, this.onTouchEnd);
    },
    onTouchStart: function onTouchStart() {
      if (!this.moved) this.tl.fadeTo(0.5, 1);
    },
    onTouchEnd: function onTouchEnd() {
      if (!this.moved) {
        this.tl.fadeTo(1, 1);
        var nodes = this.parentNode.childNodes;
        for (var j = 0; j < 4; j++) {
          if (this.moved) break;
          var x = Math.cos(j / 2 * Math.PI) | 0;
          var y = Math.sin(j / 2 * Math.PI) | 0;
          for (var i in nodes) {
            var pos = nodes[i].position;
            if (pos === this.position + x + y * 4) {
              if (nodes[i].x === SIZE) {
                nodes[i].position = this.position;
                this.position = pos;
                this.moved = true;
                this.tl.moveTo(this.position % 4 * TILE_SIZE, (this.position / 4 | 0) * TILE_SIZE, 3, QUAD_EASEOUT);
                this.tl.then(function () {
                  this.moved = false;
                  game.endCheck();
                });
                break;
              }
            }
          }
        }
      }
	  playSound('res/msg/msg.wav');
    }
  });

  game.onload = function () {
    game.rootScene.backgroundColor = '#000000';

    //パネルシャッフル
    var position = [0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15];
    var p = rand(16);

    for (var i = 0; i < 99; i++) {
      var d = rand(4);
      var x = Math.cos(d / 2 * Math.PI) | 0;
      var y = Math.sin(d / 2 * Math.PI) | 0;
      if (p % 4 === 0 && x < 0 || p % 4 === 3 && x > 0 || p < 4 && y < 0 || p > 11 && y > 0) continue;
      d = p + x + y * 4;
      var tmp = position[p];
      position[p] = position[d];
      position[d] = tmp;
      p = d;
    }

    //パネル配置
    var panel = [];
    for (var i = 0; i < 16; i++) {
      panel[i] = new Panel(TILE_SIZE, TILE_SIZE);
      panel[i].position = position[i];
      panel[i].frame = i;
      panel[i].x = panel[i].position % 4 * TILE_SIZE;
      panel[i].y = (panel[i].position / 4 | 0) * TILE_SIZE;
      game.rootScene.addChild(panel[i]);
    }
    panel[p].x = panel[p].y = SIZE;

    //グリッド表示
    var grid = new Grid();
    game.rootScene.addChild(grid);

    //ゲーム終了判定
    game.endCheck = function () {
      var c = 0;
      for (var i in panel) {
        if (panel[i].position == i) c++;
      }
      if (c === 16) {
        panel[p].tl.moveTo(panel[p].position % 4 * TILE_SIZE, (panel[p].position / 4 | 0) * TILE_SIZE, game.fps / 2, QUAD_EASEOUT);
        panel[p].tl.then(function () {
          game.rootScene.removeChild(grid);
          var endTime = new Date().getTime() - game.startTime;
          game.rootScene.removeChild(grid);
		  alert('good,you success! Time:'+(endTime/1000)+"秒");
        });
      }
    };
  };

  //ゲーム開始時間取得
  game.onstart = function () {
    game.startTime = new Date().getTime();
	 
  };

  //ゲーム開始
  game.onstart();
  // alert('good,you success! Time:'+ game.startTime );
  game.start();
  //playSound('res/msg/bac.wav');
};

