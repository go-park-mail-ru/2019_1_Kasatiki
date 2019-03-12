var canvas = document.getElementById("gameCanvas");
var ctx = canvas.getContext("2d");

var dx = 5, dy = 5,
	x = 200, y = 200;

var enemyX = 500, enemyY = 600; 


function drawBall() {
    ctx.beginPath();
    ctx.arc(x, y, 10, 0, Math.PI*2);
    ctx.fillStyle = "#008000";
    ctx.fill();
    ctx.closePath();
}

function drawEnemy() {
    ctx.beginPath();
    ctx.rect(enemyX, enemyY, 50, 50);
    ctx.fillStyle = "#0095DD";
    ctx.fill();
    ctx.closePath();
}

function gameloop() {
    ctx.clearRect(0, 0, canvas.width, canvas.height);
    drawBall();
    drawEnemy();

    if (x > enemyX && x < enemyX + 50 && y > enemyY && y < enemyY + 50 ) {
    	while (x > enemyX && x < enemyX + 50 && y > enemyY && y < enemyY + 50) {
    		enemyX = Math.random() * 650;
    		enemyY = Math.random() * 650;
    	}
		var win = window.open('http://ya.ru', '_blank');
	    var timer = setInterval(function() {
	        if (win.closed()) {
	            clearInterval(timer);
	            alert("'Secure Payment' window closed !");
	        }
	    }, 500);
    }
    else {
    	if (enemyY+25 <= y && enemyX+25 <= x) {
    		enemyY += 1; 
    		enemyX += 1;
    	}
    	else if (enemyX+25 >= x && enemyY+25 >= y) {
    		enemyY -= 1;
    		enemyX -= 1;
    	}
    	else if (enemyX+25 <= x && enemyY+25 >= y) {
    		enemyX += 1;
    		enemyY -= 1;
    	}
    	else if (enemyY+25 <= y && enemyX+25 >= x) {
    		enemyX -= 1;
    		enemyY += 1;
    	}
    }



    if(rightPressed) {
        x += 7;
    }
    if(leftPressed) {
        x -= 7;
    }
    if(upPressed) {
        y -= 7;
    }
    if(downPressed) {
        y += 7;
    }

    if (x < 0) {
    	x = canvas.width;
    }
    else if (x > canvas.width)
    	x = 0;

    if (y < 0) {
    	y = canvas.height;
    }
    else if (y > canvas.height)
    	y = 0;

}

document.addEventListener("keydown", keyDownHandler, false);
document.addEventListener("keyup", keyUpHandler, false);

function keyDownHandler(e) {
    if(e.key == "Right" || e.key == "ArrowRight") {
        rightPressed = true;
    }
    else if(e.key == "Left" || e.key == "ArrowLeft") {
        leftPressed = true;
    }
    else if(e.key == "Up" || e.key == "ArrowUp") {
        upPressed = true;
    }
    else if(e.key == "Down" || e.key == "ArrowDown") {
        downPressed = true;
    }
}

function keyUpHandler(e) {
    if(e.key == "Right" || e.key == "ArrowRight") {
        rightPressed = false;
    }
    else if(e.key == "Left" || e.key == "ArrowLeft") {
        leftPressed = false;
    }
    else if(e.key == "Up" || e.key == "ArrowUp") {
        upPressed = false;
    }
    else if(e.key == "Down" || e.key == "ArrowDown") {
        downPressed = false;
    }
}

setInterval(gameloop, 15);