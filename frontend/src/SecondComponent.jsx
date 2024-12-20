import NameContext from './NameContext';
import { useNavigate } from 'react-router-dom';
import React, { useState, useContext } from 'react';
import axios from 'axios';
import './SecondComponent.css'; // Import the CSS file

function SecondComponent() {
    const { name, setPoll } = useContext(NameContext);
    const navigate = useNavigate();

    const [pollName, setPollName] = useState('');
    const [pollId, setPollId] = useState('');
    const [joinPollId, setJoinPollId] = useState(''); // State for joining a poll
    const [pollQuestion, setPollQuestion] = useState('');
    const [pollOptions, setPollOptions] = useState(['']);
    const [loading, setLoading] = useState(false);
    const [response, setResponse] = useState(null);
    const [error, setError] = useState(null);

    const handleOptionChange = (index, value) => {
        const updatedOptions = [...pollOptions];
        updatedOptions[index] = value;
        setPollOptions(updatedOptions);
    };

    const handleAddOption = () => {
        setPollOptions([...pollOptions, '']);
    };

    const handleClick = async () => {
        setLoading(true);
        setError(null);

        try {
            const pollData = {
                name: pollName,
                question: pollQuestion,
                options: pollOptions.filter(option => option), // Filter out empty options
            };

            const result = await axios.post(`http://localhost:8080/poll`, pollData, {
                headers: {
                    'Content-Type': 'application/json',
                },
            });

            console.log(result.data);
            setPollId(result.data.id); // Assuming the ID is in result.data.id
            setResponse(result.data);
            setPoll(result.data); // Set the created poll to context

        } catch (err) {
            setError(err.message);
        } finally {
            setLoading(false);
        }
    };

    const handlePollName = async () => {
        if (!joinPollId) return; // Prevent empty ID submission

        try {
            const result = await axios.get(`http://localhost:8080/poll/${joinPollId}`, {
                headers: {
                    'Content-Type': 'application/json',
                },
            });

            setPoll(result.data);
            setResponse(result.data);
            navigate('/third');
        } catch (err) {
            setError(err.message);
        }
    };

    return (
        <div className="second-component-container">
            <h1 className="second-component-h1">Hi {name}!!</h1>

            <input
                type="text"
                placeholder="Poll Name"
                className="poll-input"
                value={pollName}
                onChange={(e) => setPollName(e.target.value)}
            />
            <input
                type="text"
                placeholder="Poll Question"
                className="poll-input"
                value={pollQuestion}
                onChange={(e) => setPollQuestion(e.target.value)}
            />

            <h3>Poll Options:</h3>
            {pollOptions.map((option, index) => (
                <div key={index}>
                    <input
                        type="text"
                        placeholder={`Option ${index + 1}`}
                        className="poll-input"
                        value={option}
                        onChange={(e) => handleOptionChange(index, e.target.value)}
                    />
                </div>
            ))}

            {/* Button container to align buttons next to each other */}
            <div className="button-container">
                <button className="add-option-button" onClick={handleAddOption}>Add Option</button>
                <button className="submit-poll-button" onClick={handleClick} disabled={loading || !pollName || !pollQuestion}>
                    {loading ? 'Sending...' : 'Submit Poll'}
                </button>
            </div>

            {response && <div className="response">{JSON.stringify(response, null, 2)}</div>}
            {error && <p className="error-message">{error}</p>}

            {/* Container for Join Poll option */}
            <div className="join-poll-container">
                <h2>Join a Poll:</h2>
                <input
                    type="text"
                    placeholder="Enter Poll ID"
                    className="poll-input"
                    value={joinPollId}
                    onChange={(e) => setJoinPollId(e.target.value)}
                />
                <button className="join-poll-button" onClick={handlePollName} disabled={!joinPollId}>
                    Join Poll
                </button>
            </div>
        </div>
    );
}

export default SecondComponent;
