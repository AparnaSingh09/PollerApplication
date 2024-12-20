import React, { useState, useContext } from "react";
import { useNavigate } from 'react-router-dom';
import NameContext from "./NameContext";
import './FirstComponent.css'; // Import the CSS file

function FirstComponent() {
    const { setName } = useContext(NameContext);
    const [inputName, setInputName] = useState('');
    const navigate = useNavigate();

    const handleEnterName = (event) => {
        setInputName(event.target.value);
        console.log(inputName);
    };

    const handleNext = () => {
        setName(inputName);
        navigate('/second');
    };

    return (
        <div className="first-component-container">
            <h1 className="first-component-h1">POLLER APP</h1>
            <input
                type="text"
                id="nameInput"
                className="input-name"
                value={inputName}
                onChange={handleEnterName}
                placeholder="Enter your name"
            />
            <button className="next-button" onClick={handleNext}>Next</button>
        </div>
    );
}

export default FirstComponent;
