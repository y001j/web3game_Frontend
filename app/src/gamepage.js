class BootScene extends Phaser.Scene {
    create ()
    {
        this.add.text(0, 0, 'Click to add new Scene');
        this.input.once('pointerdown', function () {

            this.scene.add('MainScene', MainScene, true, { x: 400, y: 300 });
        }, this);
    }
}
let gameid;
let checkGameInt;
let checkGameEndInt;
class MainScene extends Phaser.Scene {
    preload() {
        //this.load.setBaseURL('http://labs.phaser.io');
        this.load.image('stage1', 'img/game_03.png');
        this.load.image('enter', 'img/entop.png');
        this.load.image('intro', 'img/intop.png');
        this.load.image('enter2', 'img/enterl.png');
    }

    create() {
        this.cameras.main.setBackgroundColor("#fff");
        this.add.image(600, 400, 'stage1');
        let enterbutton = this.add.sprite(345, 340, 'enter');

        let enter2button = this.add.image(345, 340, 'enter2');
        enter2button.displayWidth = 300;
        enter2button.displayHeight = 160;
        enter2button.visible = false;
        let introbutton = this.add.image(860, 340, 'intro');
        const shader = this.add.shader('HSL', 400, 300, 512, 512);
        enterbutton.displayWidth = 300;
        enterbutton.displayHeight = 160;
        enterbutton.alpha = 0.6;
        introbutton.displayWidth = 300;
        introbutton.displayHeight = 160;
        introbutton.alpha = 0.6;

        enterbutton.setInteractive();
        introbutton.setInteractive();
        enterbutton.on('pointerover', function (pointer) {
            enterbutton.alpha = 1;
        }, this);
        enterbutton.on('pointerout', function (pointer) {
            enterbutton.alpha = 0.5;
        }, this);
        enterbutton.on('pointerdown', function (pointer) {
            if(window.playerWalletAddress == null){
                alert("Please Login wallet first!!");
            }
            joingame(window.playerWalletAddress).then(res => {
                if(res === true) {
                    this.scene.add('Wait', WaitPanel, true, {x: 100, y: 300});
                    checkGameInt =setInterval("checkgame()",1000);
                    checkGameEndInt = setInterval("checkgameend()",1000);
                }
                else{
                    this.scene.add('Game', GameScene, true, {x: 400, y: 300});
                    //checkGameInt =setInterval("checkgame()",1000);
                    checkGameEndInt = setInterval("checkgameend()",1000);
                    //this.scene.add('Wait', WaitPanel, true, {x: 100, y: 300});
                }
            });
        }, this);

        introbutton.on('pointerover', function (pointer) {
            introbutton.alpha = 1;
        }, this);
        introbutton.on('pointerout', function (pointer) {
            introbutton.alpha = 0.5;
        }, this);
    }
    update() {

    }
}
var timerEvent;
var timetext;
var initialTime = 150;
var showedcard = '';
var matchid;
var gameuid;
let rocknum;
let scissornum;
let papernum;
let totalscore;

let myrockynum;
let myscissornum;
let mypapernum;
let mystarnum;

let matchGameInt;
let checkMatchResultInt;

let myplayer;
let buttontext;
let buttontext2;

let showrock;
let showscissor;
let showpaper;

let bigrock;
let bigscissor;
let bigpaper;

let back;
let vs;

let matchingtext;

let brock;
let bscissor;
let bpaper;

let win;
let lose;
let tied;

class GameScene extends Phaser.Scene {

    preload() {
        //this.load.setBaseURL('http://labs.phaser.io');
        this.load.image('back', 'img/game_27.png');

        this.load.image('backgray', 'img/graybg7.png');
        this.load.image('rock', 'img/game_34.png');
        this.load.image('scissor', 'img/game_36.png');
        this.load.image('paper', 'img/game_38.png');
        this.load.image('rockcards', 'img/game_17.png');
        this.load.image('scissorcards', 'img/game_19.png');
        this.load.image('papercards', 'img/game_21.png');

        this.load.image('player', 'img/people.png');
        this.load.image('splitter', 'img/djbg_17.png');
        this.load.image('star', 'img/game_41.png');

        this.load.image('button', 'img/anl_17.png');
        //对手玩家相关的东西
        //this.load.image('player', 'img/people.png');
        this.load.image('rockcardsgr', 'img/game_52.png');
        this.load.image('scissorcardsgr', 'img/game_49.png');
        this.load.image('papercardsgr', 'img/game_54.png');
        this.load.image('blankcardsgr', 'img/game_bg.png');

        this.load.image('vspic', 'img/vs.png');

        this.load.image('brock', 'img/d3.png');
        this.load.image('bpaper', 'img/y2.png');
        this.load.image('bscissor', 'img/y3.png');

        this.load.image('win', 'img/win.png');
        this.load.image('lose', 'img/lose.png');
        this.load.image('tied', 'img/tied.png');
    };

    create() {
        this.cameras.main.setBackgroundColor("#fff");
        back = this.add.image(600, 400, 'back');
        vs = this.add.image(600, 360, 'vspic').setScale(0.7);
        vs.visible = false;
        //back.displayWidth = 600;
        //back.displayHeight = 220;
        this.add.image(500, 100, 'rockcards');
        this.add.image(600, 100, 'scissorcards');
        this.add.image(700, 100, 'papercards');

        myplayer = this.add.image(280, 350, 'player').setScale(0.5);
        myplayer.visible = false;
        rocknum = this.add.text(480, 150, '×80', { font: "22px Arial", fill: "#000" });
        scissornum = this.add.text(580, 150, '×80', { font: "22px Arial", fill: "#000" });
        papernum = this.add.text(680, 150, '×80', { font: "22px Arial", fill: "#000" });

        //this.add.text(100, 150, 'COUNT DOWN', { font: "22px Arial", fill: "#000" });
        totalscore = this.add.text(1000, 150, 'SCORE', { font: "22px Arial", fill: "#000" });

        let splitter = this.add.image(600, 600, 'splitter');
        splitter.displayWidth = 1200;

        showrock = this.add.image(100, 700, 'rock');
        brock = this.add.image(280, 360, 'brock').setScale(0.5);
        brock.visible = false;
        showrock.on('pointerover', function (pointer) {
            showrock.setScale(1.1);
        }, this);
        showrock.on('pointerout', function (pointer) {
            showrock.setScale(1);
        }, this);
        showrock.on('pointerdown', function (pointer) {
            brock.visible = true;
            bpaper.visible = false;
            bscissor.visible = false;
            buttontext2.setInteractive();
            showedcard = "rock";
        }, this);
        showscissor = this.add.image(200, 700, 'scissor');
        bscissor = this.add.image(280, 360, 'bscissor').setScale(0.5);
        bscissor.visible = false;
        showscissor.on('pointerover', function (pointer) {
            showscissor.setScale(1.1);
        }, this);
        showscissor.on('pointerout', function (pointer) {
            showscissor.setScale(1);
        }, this);
        showscissor.on('pointerdown', function (pointer) {
            brock.visible = false;
            bpaper.visible = false;
            bscissor.visible = true;
            buttontext2.setInteractive();
            showedcard = "scissor";
        }, this);
        showpaper = this.add.image(300, 700, 'paper');
        bpaper = this.add.image(280, 360, 'bpaper').setScale(0.5);
        bpaper.visible = false;
        showpaper.on('pointerover', function (pointer) {
            showpaper.setScale(1.1);
        }, this);
        showpaper.on('pointerout', function (pointer) {
            showpaper.setScale(1);
        }, this);
        showpaper.on('pointerdown', function (pointer) {
            brock.visible = false;
            bpaper.visible = true;
            bscissor.visible = false;
            buttontext2.setInteractive();
            showedcard = "paper";
        }, this);
        myrockynum = this.add.text(90, 760, '×', { font: "26px Arial Black", fill: "#000" });
        myscissornum = this.add.text(190, 760, '×', { font: "26px Arial Black", fill: "#000" });
        mypapernum = this.add.text(290, 760, '×', { font: "26px Arial Black", fill: "#000" });
        mystarnum = this.add.text(1000, 760, '×', { font: "26px Arial Black", fill: "#000" });

        this.add.image(940, 700, 'star').setScale(0.7);
        this.add.image(1000, 700, 'star').setScale(0.7);
        this.add.image(1060, 700, 'star').setScale(0.7);

        //myplayer = this.add.image()

        //player.displayWidth = 180;
       //player.displayHeight = 200;

        let cbutton = this.add.sprite(600, 600, 'button');
        cbutton.setInteractive();
        buttontext = this.add.text(540, 575, 'PLAY', { font: "40px Arial Black", fill: "#FFC90E" });
        buttontext2 = this.add.text(518, 575, 'CHECK', { font: "40px Arial Black", fill: "#FFC90E" });
        buttontext2.visible = false;
        buttontext2.alpha = 0.7;
        buttontext.setInteractive();
        buttontext.alpha = 0.6;
        matchingtext = this.add.text(430, 502, 'Matching Game!!', { font: "40px Arial Black", fill: "#FFC90E" });
        matchingtext.visible = false;
        buttontext.on('pointerover', function (pointer) {
            buttontext.alpha = 1;
            buttontext.setScale(1.05);
        }, this);
        buttontext.on('pointerout', function (pointer) {
            buttontext.alpha = 0.6;
            buttontext.setScale(1);
        }, this);
        buttontext.on('pointerdown', function (pointer) {
            buttontext.disableInteractive();
            myplayer.visible = true;
            back.setScale(0.7);
            back.alpha = 0.7;
            win.visible = false;
            lose.visible = false;
            tied.visible = false;
            brock.visible = false;
            bpaper.visible = false;
            bscissor.visible = false;
            matchgame(window.playerWalletAddress).then(res => {
                    //this.scene.add('Wait', WaitPanel, true, {x: 100, y: 300});
                matchingtext.visible = true;
                matchGameInt=setInterval("checkGameforMatch()",1000);
                    //this.scene.add('Wait', WaitPanel, true, {x: 100, y: 300});

            });
        }, this);

        win = this.add.image(200, 400, 'win').setScale(1);
        win.visible = false;
        lose = this.add.image(200, 400, 'lose').setScale(1);
        lose.visible = false;
        tied = this.add.image(200, 400, 'tied').setScale(1);
        tied.visible = false;

        buttontext2.on('pointerover', function (pointer) {
            buttontext2.alpha = 1;
            buttontext2.setScale(1.05);
        }, this);
        buttontext2.on('pointerout', function (pointer) {
            buttontext2.alpha = 0.7;
            buttontext2.setScale(1);
        }, this);
        buttontext2.on('pointerdown', function (pointer) {
            if(showedcard == ''){
                alert('Please choose a card!');
                return
            }
            buttontext2.disableInteractive();

            showcard(showedcard).then(res => {
                showrock.disableInteractive();
                showpaper.disableInteractive();
                showscissor.disableInteractive();
                checkMatchResultInt = setInterval("checkMatchResult()", 1000);
            });
        }, this);


        timetext = this.add.text(100, 150, 'COUNTDOWN: ' , { font: "22px Arial", fill: "#000" });
        // Each 1000 ms call onEvent
        this.time.addEvent({ delay: 1000, callback: this.onEvent, callbackScope: this, loop: true });

        //this.scene.add('contest', Contest, true);
        //timerEvent = this.time.addEvent({ delay: 1000, repeat:1 });
    }
    update() {
        //Countdown(timerEvent);
    }
    onEvent () {
        let thedata;
        axios.get('http://localhost:8081/getgame', {
            params: {
                gameid: gameid,
                addr: window.playerWalletAddress + ''
            }
        }).then(function (response) {
            thedata = eval(response.data);
            rocknum.setText('×' + thedata.Total.Rock);
            scissornum.setText('×' + thedata.Total.Scissor);
            papernum.setText('×' + thedata.Total.Paper);
            totalscore.setText('SCORE:' + thedata.Total.Score);
            //console.log(thedata.Members);
            //console.log(window.playerWalletAddress);
            myrockynum.setText('×' + thedata.Members['0'].Cards.Rock);
            if(thedata.Members['0'].Cards.Rock == 0){
                showrock.disableInteractive();
                showrock.alpha = 0.5;
            }
            myscissornum.setText('×' + thedata.Members['0'].Cards.Scissor);
            if(thedata.Members['0'].Cards.Scissor == 0){
                showscissor.disableInteractive();
                showscissor.alpha = 0.5;
            }
            mypapernum.setText('×' + thedata.Members['0'].Cards.Paper);
            if(thedata.Members['0'].Cards.Paper == 0){
                showpaper.disableInteractive();
                showpaper.alpha = 0.5;
            }
            mystarnum.setText('×' + thedata.Members['0'].Cards.Star);
            if(thedata.Members['0'].Cards.Star == 0 ||(thedata.Members['0'].Cards.Rock == 0&& thedata.Members['0'].Cards.Paper == 0 && thedata.Members['0'].Cards.Scissor == 0)){
                buttontext.visible = false;
                buttontext2.visible = false;
                //game.scene.add('WaitEnd', WaitPanelEnd, true, {x: 400, y: 300});
            }
            var endtime = new Date(thedata.end);
            var nowtime = new Date();
            var total = parseInt((endtime - nowtime)/1000);
            if(total >=0){
                timetext.setText('COUNTDOWN: ' + formatTime(total))
            }
        }).catch(function (error) {
            // handle error
            console.log(error);
        })
    };
    //update my cards

}

class WaitPanel extends Phaser.Scene {
    create() {
        this.cameras.main.setBackgroundColor("#fff");
        const text2 = this.add.text(100, 200, 'Please Wait Other Players Joining...', { font: '64px Arial' });
        text2.setTint(0xff00ff, 0xffff00, 0x0000ff, 0xff0000);
    }
}

class WaitPanelEnd extends Phaser.Scene {
    create() {
        this.cameras.main.setBackgroundColor("#fff");
        const text3 = this.add.text(100, 200, 'Please Wait Game End...', { font: '64px Arial' });
        text3.setTint(0xff00ff, 0xffff00, 0x0000ff, 0xff0000);
    }
}

let conteststar;
let contestrock;
let contestscissor;
let contestpaper;
let contestblank;
//对手显示的内容
var Contest = new Phaser.Class({
    Extends: Phaser.Scene,
    initialize:
        function Contest()
        {
            Phaser.Scene.call(this);
        },
    create: function ()
    {
        player = this.add.image(880, 350, 'player').setScale(0.5);
        this.add.image(1020, 280, 'rockcardsgr');
        this.add.image(1020, 350, 'scissorcardsgr');
        this.add.image(1020, 350, 'blankcardsgr');
        this.add.image(1020, 420, 'papercardsgr');
        smallstart = this.add.image(880, 480, 'star');
        smallstart.displayWidth = 30;
        smallstart.displayHeight = 30;
        conteststar = this.add.text(900, 470, '', { font: "24px Arial", fill: "#000" });
        contestrock =  this.add.text(1050, 270, '', { font: "24px Arial", fill: "#000" });
        contestscissor = this.add.text(1050, 340, '', { font: "24px Arial", fill: "#000" });
        contestblank = this.add.text(1080, 340, '', { font: "24px Arial", fill: "#000" });
        contestblank.visible = false;
        contestpaper = this.add.text(1050, 410, '', { font: "24px Arial", fill: "#000" });
    },
});

var ScoreBoard = new Phaser.Class({
    Extends: Phaser.Scene,
    initialize:
        function Contest()
        {
            Phaser.Scene.call(this);
        },
    create: async function ()
    {
        this.cameras.main.setBackgroundColor("#fff");
        const text2 = this.add.text(480, 40, 'Score Board', { font: '48px Arial' });
        text2.setTint(0xff00ff, 0xffff00, 0x0000ff, 0xff0000);

        exitbuttion = this.add.text(960, 720, '🚶Exit', { font: '48px Arial',fill: "#000" });
        exitbuttion.setInteractive();
        exitbuttion.on('pointerdown', function (pointer) {
             axios.get('http://localhost:8081/removeplayer', {
                params: {
                    addr: window.playerWalletAddress
                }
            }).then(function (response) {
                //thedata = eval(response.data);
            })
                .catch(function (error) {
                    // handle error
                    console.log(error);
                })
            location.reload();
        }, this);

        let thedata;
        await axios.get('http://localhost:8081/getscores', {
            params: {
                gameuid: gameuid
            }
        }).then(function (response) {
            thedata = eval(response.data);
        })
            .catch(function (error) {
                // handle error
                console.log(error);
            })
        console.log(gameuid);
        console.log(thedata);
        var h = 0;
        for(var item of thedata) {
            if(item.address.substring(0,10) === window.playerWalletAddress.substring(0,10)) {
                textb = this.add.text(160, 100 + h, item.score + '  ' + item.bonus + '  ' + item.address.substring(0,20) + '...', {
                    font: '30px Courier'
                });
                textb.setTint(0xff00ff, 0xffff00, 0x0000ff, 0xff0000);
            }
            else {
                this.add.text(160, 100 + h, item.score + '  ' + item.bonus + '  ' + item.address.substring(0,20)  + '...', {
                    font: '30px Courier',
                    fill: "#000"
                });
            }
            h = h+40;
        }
    },
});

function formatTime(seconds){
    // Minutes
    var minutes = Math.floor(seconds/60);
    // Seconds
    var partInSeconds = seconds%60;
    // Adds left zeros to seconds
    partInSeconds = partInSeconds.toString().padStart(2,'0');
    // Returns formated time
    return `${minutes}:${partInSeconds}`;
}
async function joingame(addr){
    let thedata;
    await axios.get('http://localhost:8081/joingame', {
        params: {
            addr: addr
        }
    }).then(function (response) {
        thedata = eval(response.data);
    })
        .catch(function (error) {
            // handle error
            console.log(error);
        })
    console.log(thedata);
    gameid = thedata.gameid;
    if (thedata.message === "success"){
        return true;
    }
    else {
        return false;
    }
}

async function matchgame(addr){
    let thedata;
    await axios.get('http://localhost:8081/playermatch', {
        params: {
            playerid: addr,
            gameid: gameid
        }
    }).then(function (response) {
        thedata = eval(response.data);
    })
        .catch(function (error) {
            // handle error
            console.log(error);
        })
    console.log(thedata);
    //gameid = thedata.gameid;
}

//玩家出牌
async function showcard(card){
    let thedata;
    await axios.get('http://localhost:8081/playershowcard', {
        params: {
            playerid: window.playerWalletAddress,
            gameid: gameid,
            card:card
        }
    }).then(function (response) {
        thedata = eval(response.data);
        matchid = thedata.matchid;
    })
        .catch(function (error) {
            // handle error
            console.log(error);
        })
    console.log(thedata);
}

async function checkgame(){
    let thedata;
    await axios.get('http://localhost:8081/getgame', {
        params: {
            gameid: gameid
        }
    }).then(function (response) {
        thedata = eval(response.data);
    }).catch(function (error) {
        // handle error
        console.log(error);
    })
    console.log(thedata);
    if (thedata.status === 0){
        return true;
    }
    if(thedata.status === 1){
        //游戏状态为1，也就是进行中，则立马进入游戏
        //clearInterval(checkGameInt);
        clearInterval(checkGameInt);
        game.scene.remove("Wait");
        game.scene.add('Game', GameScene, true, {x: 400, y: 300});
    }
    // if(thedata.status === 2){ //如果游戏状态为2，则游戏结束
    //     clearInterval(checkGameInt);
    //     gameuid = thedata.id;
    //     //game.scene.remove("Game");
    //     game.scene.add('Score', ScoreBoard, true, {x: 400, y: 300});
    // }
}

async function checkgameend(){
    let thedata;
    await axios.get('http://localhost:8081/getgame', {
        params: {
            gameid: gameid
        }
    }).then(function (response) {
        thedata = eval(response.data);
    }).catch(function (error) {
        // handle error
        console.log(error);
    })
    if(thedata.status === 2){ //如果游戏状态为2，则游戏结束
        clearInterval(checkGameEndInt);
        gameuid = thedata.id;
        //game.scene.remove("Game");
        game.scene.add('Score', ScoreBoard, true, {x: 400, y: 300});
    }
}


async function checkGameforMatch(){
    let thedata;
    await axios.get('http://localhost:8081/getgame', {
        params: {
            gameid: gameid,
            addr: window.playerWalletAddress
        }
    }).then(function (response) {
        thedata = eval(response.data);
    }).catch(function (error) {
        console.log(error);
    })
    console.log(thedata);
    if (thedata.Members['0'].status === 2){
        buttontext.visible = false;
        matchingtext.visible = false;
        game.scene.add('contest', Contest, true);
        buttontext2.visible = true;
        buttontext2.setInteractive();
        showrock.setInteractive();
        showpaper.setInteractive();
        showscissor.setInteractive();
        clearInterval(matchGameInt);
        //显示出对手的信息

        contestpaper.setText('×' + thedata.Members['1'].Cards.Paper);
        if(thedata.Members['1'].Cards.Rock == 0||thedata.Members['1'].Cards.Scissor == 0|| thedata.Members['1'].Cards.Paper == 0) {
            thetotal = thedata.Members['1'].Cards.Rock + thedata.Members['1'].Cards.Scissor + thedata.Members['1'].Cards.Paper
            contestblank.setText('Totol x' + thetotal);
            contestblank.visible = true;
            contestrock.visible = false;
            contestscissor.visible = false;
            contestpaper.visible = false;
        } else{
            conteststar.setText('×' + thedata.Members['1'].Cards.Star);
            contestrock.setText('×' + thedata.Members['1'].Cards.Rock);
            contestscissor.setText('×' + thedata.Members['1'].Cards.Scissor);
            contestrock.visible = true;
            contestscissor.visible = true;
            contestpaper.visible = true;
        }

        vs.visible = true;
        //return true;
    }
}

async function checkMatchResult(){
    let thedata;
    await axios.get('http://localhost:8081/getmatch', {
        params: {
            gameid: gameid,
            matchid: matchid
        }
    }).then(function (response) {
        thedata = eval(response.data);
    }).catch(function (error) {
        console.log(error);
    })
    //测试阶段显示信息
    console.log(thedata);
    if (thedata.result != 0){
        buttontext.visible = true;
        buttontext.setInteractive();
        matchingtext.visible = false;
        game.scene.remove('contest');
        buttontext2.visible = false;
        buttontext2.disableInteractive();
        showrock.disableInteractive();
        showpaper.disableInteractive();
        showscissor.disableInteractive();
        clearInterval(checkMatchResultInt);
        showedcard = '';

        brock.visible = false;
        bpaper.visible = false;
        bscissor.visible = false;
        //显示出对手的信息
        if(thedata.result == 1 && thedata.Player1Addr == window.playerWalletAddress) {
            win.visible = true;
        } else if (thedata.result == 1 && thedata.Player1Addr != window.playerWalletAddress){
            lose.visible = true;
        }
        if(thedata.result == 2 && thedata.Player1Addr == window.playerWalletAddress) {
            lose.visible = true;
        } else if (thedata.result == 2 && thedata.Player1Addr != window.playerWalletAddress){
            win.visible = true;
        }
        if(thedata.result == 3) {
            tied.visible = true;
        }

        vs.visible = false;
        //return true;
    }
}
//var Phaser = require("phaser")
var config = {
    type: Phaser.AUTO,
    width: 1200,
    height: 800,
    parent: 'main',
    physics: {
        default: 'arcade',
        arcade: {
            gravity: { y: 200 }
        }
    },
    scene:MainScene
};
var game = new Phaser.Game(config);
