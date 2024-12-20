//Update ARRAYS in state

import React,{useState} from "react";


function Food() {
    const [foods, setFoods] = useState(['Apple','Banana','Pears'])


    const handleAddFood = () => {
        const newFood = document.getElementById("foodInput").value
        //reset back to ""
        document.getElementById("foodInput").value = ""
        //setFoods(...foods, newFood)
        setFoods(f => [...f,newFood])
    }

    return (<div>
        <h2>List of foods</h2>
        <ul>
            {foods.map((food, index) => <li key={index}>{food}</li>)}
        </ul>
        <input type="text" id="foodInput" placeholder="Enter food name"></input>
        <button onClick={handleAddFood}>Add Food</button>
    </div>)

}

export default Food