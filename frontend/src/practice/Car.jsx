import { useState } from "react"

//UPdate state of an object using react
function Car() {
    const [car, setCar] = useState({year: 2024,
                                    make: "Ford",
                                    model: "Mustang"});

        const handleYearChange = (event) => {
            setCar({...car, year: event.target.value})
        }

        const handleMakeChange = (event) => {
            setCar({...car, make: event.target.value})
        }

        const handleModelChange = (event) => {
            setCar({...car, model: event.target.value})
        }

        return(<div>
            <p>Your fav car is: {car.year} {car.make} {car.model}</p>
            <input type="number" value={car.year} onChange={handleYearChange}></input><br></br>
            <input type="text" value={car.make} onChange={handleMakeChange}></input><br></br>
            <input type="text" value={car.model} onChange={handleModelChange}></input><br></br>

        </div>)  

}

export default Car