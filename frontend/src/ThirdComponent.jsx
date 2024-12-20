import NameContext from "./NameContext";
import React, { useState, useContext, useEffect } from 'react';
import axios from 'axios';
import './ThirdComponent.css'; // Import the CSS file

function ThirdComponent() {
    console.log("ThirdComponent");
    const { poll } = useContext(NameContext); 
    console.log("Poll data:", poll);
    
    const [selectedOption, setSelectedOption] = useState('');
    const [updatedPoll, setUpdatedPoll] = useState(null);
    const [isConfirmed, setIsConfirmed] = useState(false); // Track confirmation

    // Fetch poll function defined inside the component
    const fetchPoll = async () => {
        if (poll && poll.id) {
            try {
                const response = await axios.get(`http://localhost:8080/poll/${poll.id}`, {
                    headers: {
                        'Content-Type': 'application/json',
                    },
                });
                setUpdatedPoll(response.data);
            } catch (err) {
                console.error('Error fetching updated poll:', err.message);
            }
        }
    };

    useEffect(() => {
        fetchPoll(); // Initial fetch
           
        // Set up interval to fetch poll every 5 seconds
        const intervalId = setInterval(fetchPoll, 5000);
        console.log("fetching the data..")
        // Clear the interval on component unmount
        return () => clearInterval(intervalId);
    }, [poll]);

    if (!updatedPoll) {
        return <p>No data available.</p>;
    }

    const { name, question, options, result } = updatedPoll;

    const handleRadioChange = (event) => {
        const { value } = event.target;
        setSelectedOption(value);
    };

    const handleConfirm = async () => {
        if (!selectedOption) return; // Prevent if no option is selected

        console.log(`Selected option: ${selectedOption}`);

        try {
            await axios.get(`http://localhost:8080/updatePoll`, {
                params: {
                    id: poll.id,
                    option: selectedOption,
                },
                headers: {
                    'Content-Type': 'application/json',
                },
            });
            // Fetch the updated poll data after submitting the selected option
            await fetchPoll();
            setIsConfirmed(true); // Mark as confirmed
        } catch (err) {
            console.error('Error submitting selected option:', err.message);
        }
    };

    return (
        <div className="third-component-container">
            <h1 className="third-component-h1">{name}</h1>
            <p className="question"><strong>Question:</strong> {question}</p>

            <div className="options">
                <h2>Options:</h2>
                {options.map((option, index) => (
                    <div key={index} className="option-label">
                        <input
                            type="radio"
                            name="pollOptions" // Grouping for radio buttons
                            value={option}
                            checked={selectedOption === option}
                            onChange={handleRadioChange}
                            className="option-input"
                            disabled={isConfirmed} // Disable if confirmed
                        />
                        <label>{option}</label>
                    </div>
                ))}
            </div>

            <button 
                onClick={handleConfirm} 
                disabled={!selectedOption || isConfirmed} // Disable if no selection or already confirmed
            >
                Confirm Selection
            </button>

            <div className="results">
                <h2>Results:</h2>
                <ul>
                    {Object.entries(result).map(([option, count], index) => (
                        <li key={index}>
                            <strong>{option}:</strong> {count}
                        </li>
                    ))}
                </ul>
            </div>
        </div>
    );
}

export default ThirdComponent;
