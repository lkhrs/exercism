<?php

class PizzaPi
{
    public function calculateDoughRequirement($pizzas, $people)
    {
        return $pizzas * (($people * 20) + 200);
    }

    public function calculateSauceRequirement($pizzas, $sauceVolume)
    {
        return $pizzas * 125 / $sauceVolume;
    }

    public function calculateCheeseCubeCoverage($cheese_dimension, $cheese_thickness, $diameter)
    {
        return floor(($cheese_dimension ** 3) / ($cheese_thickness * pi() * $diameter));
    }

    public function calculateLeftOverSlices($pizzas, $friends)
    {
        return ($pizzas * 8) % $friends;
    }
}
