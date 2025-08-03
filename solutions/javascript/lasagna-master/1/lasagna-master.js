/// <reference path="./global.d.ts" />
// @ts-check

/**
 * Implement the functions needed to solve the exercise here.
 * Do not forget to export them so they are available for the
 * tests. Here an example of the syntax as reminder:
 *
 * export function yourFunction(...) {
 *   ...
 * }
 */

/**
 * Returns status messages based on the time remaining on the timer.
 * @param  {number} timer  Remaining time
 * @return {string} status message
 */
export function cookingStatus (timer) {
	if (timer === 0) {
		return 'Lasagna is done.'
	}
	if (timer > 0) {
		return 'Not done, please wait.'
	}
	return 'You forgot to set the timer.'
}

/**
 * Returns the preparation time based on the number of layers (if provided).
 * @param  {array} layers  An array of layers
 * @param  {number} time   Prep time per layer
 * @return {number}        Total prep time
 */
export function preparationTime(layers, time) {
	if (time) {
		return layers.length * time
	}
	return layers.length * 2
}

/**
 * Takes an array of layers and returns the total amount of noodles and sauce needed.
 * @param  {array} layers Array of layers.
 * @return {object}        Object with noodle and sauce quantities
 */
export function quantities(layers) {
	let noodles = layers.filter(layer => layer === 'noodles').length * 50
	let sauce = layers.filter(layer => layer === 'sauce').length * 0.2
	return {
		"noodles": noodles,
		"sauce": sauce
	}
}

/**
 * Takes the last ingredient from a list and adds it to another list.
 * @param {array} friendsList list of ingredients from friend
 * @param {array} myList      your list of ingredients
 */
export function addSecretIngredient(friendsList, myList) {
	myList.push(friendsList.slice(-1)[0])
	return
}

/**
 * Scales a recipe for a number of portions.
 * @param  {object} recipe   Recipe to scale.
 * @param  {number} portions Number of portions.
 * @return {object}          Scaled recipe.
 */
export function scaleRecipe(recipe, portions) {
	let scaledRecipe = Object.create(recipe)
	for (let key in scaledRecipe) {
		scaledRecipe[key] = scaledRecipe[key] / 2 * portions
	}
	return scaledRecipe
}
