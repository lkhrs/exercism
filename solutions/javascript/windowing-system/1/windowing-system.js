// @ts-check

/**
 * Implement the classes etc. that are needed to solve the
 * exercise in this file. Do not forget to export the entities
 * you defined so they are available for the tests.
 */

export function Size(width = 80, height = 60) {
	this.width = width
	this.height = height
}

Size.prototype.resize = function (newWidth, newHeight) {
	this.width = newWidth
	this.height = newHeight
}

export class Position {
	constructor(x = 0, y = 0) {
		this.x = x
		this.y = y
	}

	move(newX, newY) {
		this.x = newX
		this.y = newY
	}
}

export class ProgramWindow {
	constructor() {
		this.screenSize = new Size(800, 600)
		this.size = new Size()
		this.position = new Position()
	}

	resize(newSize) {
		if (newSize.width <= 1) {
			newSize.width = 1
		}
		if (newSize.height <=1) {
			newSize.height = 1
		}
		if (newSize.width >= this.screenSize.width) {
			newSize.width = this.screenSize.width - this.position.x
		}
		if (newSize.height >= this.screenSize.height) {
			newSize.height = this.screenSize.height - this.position.y
		}
		this.size.width = newSize.width
		this.size.height = newSize.height
	}

	move(newPos) {
		if (newPos.x <= 0) {
			newPos.x = 0
		}
		if (newPos.y <= 0) {
			newPos.y = 0
		}
		if (newPos.x + this.size.width >= this.screenSize.width) {
			newPos.x = this.screenSize.width - this.size.width
		}
		if (newPos.y + this.size.height >= this.screenSize.height) {
			newPos.y = this.screenSize.height - this.size.height
		}
		this.position.x = newPos.x
		this.position.y = newPos.y
	}
}

export function changeWindow(programWindow) {
	programWindow.size.width = 400
	programWindow.size.height = 300
	programWindow.position.x = 100
	programWindow.position.y = 150
	return programWindow
}
