import React, {useState} from "react"

function ColorPicker () {
    const[color, setColor] = useState("#FFFFFF");

    const changeColor = (event) => {
        setColor(event.target.value)

    }

    return (
        <div className="color-picker-container">
            <hi>Color Picker</hi>
            <div className="color-display" style={{backgroundColor: color}}>
                <p> Selected color: {color}</p>
            </div>
            <label>Select a color:</label>
            <input type="color" value={color} onChange={changeColor}></input>
        </div>
    )

}

export default ColorPicker