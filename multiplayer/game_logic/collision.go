package game_logic

func IsCollision(obj1, obj2 *DynamycObject) bool {
	if obj1.Y > obj2.Y - obj1.Ysize && 			// граница	сверху
		obj1.Y < obj2.Y + obj2.Ysize && 		// граница	снизу
		obj1.X > obj2.X - obj1.Xsize && 		// граница	справа
		obj1.X < obj2.X + obj2.Xsize { 			// граница 	слева
		//fmt.Println("Colision with Obj_1: ", obj1.Name, " and  Obj_2: ", obj2.Name)
		return true
	}
	return false
}

func SimpleCollisionEvent(obj1, obj2 *DynamycObject, moves Moves) {
	if moves.Up || moves.Down {
		if obj1.Y-(obj2.Y-obj2.Ysize) > obj2.Y+obj2.Ysize-obj1.Y && (obj1.X < obj2.X-obj1.Xsize && obj1.X > obj2.X+obj2.Xsize) {
			obj1.Y = obj2.Y + obj2.Ysize
		} else {
			obj1.Y = obj2.Y - obj1.Ysize
		}
	}
	if moves.Left || moves.Right {
		if obj1.X-(obj2.X-obj2.Xsize) > obj2.X+obj2.Xsize-obj1.X {
			obj1.X = obj2.X + obj2.Xsize
		} else {
			obj1.X = obj2.X - obj1.Xsize
		}
	}
}

func ObjectFilter(slice []*DynamycObject) []*DynamycObject {
	keys := make(map[*DynamycObject]bool)
	list := []*DynamycObject{}
	for _, entry := range slice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

