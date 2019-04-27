export default class CollisionHandler {
    /**
     * Функция возвращает массив пар {first: firstObj, second: secondObj}
     * @param {Array of } firstArr 
     * @param {Array of } secondArr 
     */
    _getPairCollisions(firstArr, secondArr) {
        let pair;
        let pairs = [];
        firstArr.forEach((firstObj, fidx) => {
            if (typeof firstObj !== 'undefined') {
                secondArr.forEach((secondObj, sidx) => {
                    if (typeof secondObj !== 'undefined') {
                        if (this._checkCollision(firstObj, secondObj)) {
                            pair = {
                                first: {firstObj, fidx},
                                second: {secondObj, sidx}
                            };
                            pairs.push(pair);
                        }
                    }
                });
            }
        });
        return pairs;
    }

    handleCollisions(objects, score = {}) {
        let PlayerBarrierCollision = this._getPairCollisions(objects['players'], objects['barriers']);
        let BulletBarrierCollision = this._getPairCollisions(objects['bullets'], objects['barriers']);
        let BulletAdvsCollision = this._getPairCollisions(objects['bullets'], objects['advs']);
        let PlayersAdvsCollision = this._getPairCollisions(objects['players'], objects['advs']);
        let BarriersAdvsCollision = this._getPairCollisions(objects['barriers'], objects['advs']);
        let PlayersShopsCollision = this._getPairCollisions(objects['players'], objects['shops']);
        let PlayersBuffsCollision = this._getPairCollisions(objects['players'], objects['buffers']);
        let BuffsBarriersCollision = this._getPairCollisions(objects['buffers'], objects['barriers']);

        if (PlayerBarrierCollision.length != 0) {
            PlayerBarrierCollision.forEach( pair => {
                pair.first.firstObj.interact(pair.second.secondObj);
                pair.second.secondObj.interact(pair.first.firstObj);
            });
        }

        if (BulletBarrierCollision.length != 0) {
            BulletBarrierCollision.forEach( pair =>{
                pair.first.firstObj.interact();
                objects['bullets'].splice(pair.first.fidx, 1);
                pair.second.secondObj.interact();
            });

        }

        if (BulletAdvsCollision.length != 0) {
            BulletAdvsCollision.forEach( pair => {
                objects['bullets'].splice(pair.first.fidx, 1);
                let advHealth = pair.second.secondObj.interact(pair.first.firstObj);
                if (advHealth <= 0) {
                    objects['players'][0].score += 100;
                    objects['advs'].splice(pair.second.sidx, 1);
                }            
            });
        }


        if (PlayersAdvsCollision.length != 0) {
            PlayersAdvsCollision.forEach(pair => {
                pair.first.firstObj.interact('adv');
                let advHealth = pair.second.secondObj.interact(pair.first.firstObj);
                if (advHealth <= 0) {
                    objects['players'][0].score += 100;
                    objects['advs'].splice(pair.second.sidx, 1);
                }     
            })           
        }

        if (BarriersAdvsCollision.length != 0) {
            BarriersAdvsCollision.forEach( pair => {
                pair.first.firstObj.interact();
                pair.second.secondObj.interact(pair.first.firstObj);      
            })      
        }

        if (PlayersShopsCollision.length != 0) {
            PlayersShopsCollision.forEach(pair => {
                pair.first.firstObj.interact(pair.second.secondObj);
                pair.second.secondObj.player = pair.first.firstObj;
                pair.second.secondObj.playerInShop = true;   
            })
        } else {
            if (objects['shops'].length != 0) {
                objects['shops'][0].playerInShop = false;
                objects['players'][0].inShop = false;
                objects['shops'][0].close(); 
            }
        }

        if (PlayersBuffsCollision.length != 0) {
            PlayersBuffsCollision.forEach( pair => {
                pair.second.secondObj.interact(pair.first.firstObj);
                objects['buffers'].splice(pair.second.sidx, 1);
            });
        }

        if (BuffsBarriersCollision.length != 0) {
            BuffsBarriersCollision.forEach(pair => {
                objects['buffers'].splice(pair.first.fidx, 1);
            })
        }
    }


    /**
     * 
     * @param {Array of Array} map 
     * @param {Array} objArray 
     */
    _setObjectsVisibilityArea(map, objArray) {

    }
 
    _checkCollision(obj1, obj2) {
        return this._checkCollisionRectangles(obj1, obj2);
    }

    _checkCollisionRectangles(obj1, obj2) {
        // let x1 = obj1.xPos;
        // let y1 = obj1.yPos;
        // let xd1 = obj1.xPos + obj1.xSize;
        // let yd1 = obj1.yPos + obj1.ySize;

        // let x2 = obj2.xPos;
        // let y2 = obj2.yPos;
        // let xd2 = obj2.xPos + obj2.xSize;
        // let yd2 = obj2.yPos + obj2.ySize;

        // let left = Math.min(x1, x2);
        // let right = Math.max(xd1, xd2);
        // let top = Math.min(y1, y2);
        // let bottom = Math.max(yd1, yd2);

        // let width = right - left;
        // let height = bottom - top;

        // if (width <= obj1.xSize + obj2.xSize && height <= obj1.ySize + obj2.ySize) {
        //     return true;
        // }

        // return false;

        if (obj1.xPos < obj2.xPos + obj2.xSize &&
            obj1.xPos + obj1.xSize > obj2.xPos &&
            obj1.yPos < obj2.yPos + obj2.ySize &&
            obj1.yPos + obj1.ySize > obj2.yPos) {
        
            return true;
        }

        return false;

    }

    _checkCollisionCirles(obj1, obj2) {

    }

    _checkCollisionRectangleCirle(rect, cirle) {

    }
}

class increaseHpByTime {
    constructor() {
        this.start = Date.now();
    }

    start() {
        this.start = Date.now();
    }

    end(person) {
        person.hp = person.hp + Math.trunc((Date.now() - this.start) / 100);
        if (person.hp > person.hpCapacity) {
            person.hp = person.hpCapacity;
        }
    }
}