export default class Bulelt {
    constructor(
        xPos,
        yPos,
        xSize = 3,
        ySize = 3,
        color = 'red',
        // URL = "/default_texture",
        velocity = 100,
        damage,
        xDestination,
        yDestination,
    ) {
        this.xSize = xSize; // vh
        this.ySize = ySize; // vh

        this.xPos = xPos;
        this.yPos = yPos;

        this.damage = damage;
        this.color = color;

        this.name = 'bullet'

        this.texture = URL; // URL 
        this.velocity = velocity; // у.е

        this.xDestination = xDestination;
        this.yDestination = yDestination;

        this.teta = Math.atan2(this.xDestination - this.xPos, this.yDestination - this.yPos);

        console.log(this.xDestination, this.yDestination);
    }

    render(ctx) {
        ctx.beginPath();
        ctx.rect(this.xPos, this.yPos, this.xSize, this.ySize);
        ctx.fillStyle = this.color;
        ctx.fill();
        ctx.closePath();
    }

    go () {
        this.xPos += this.velocity * Math.sin(this.teta);
        this.yPos += this.velocity * Math.cos(this.teta);
    }

    interact() {
        
    }
}