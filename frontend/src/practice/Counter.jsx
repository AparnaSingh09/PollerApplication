import React, {useState} from "react"
function Counter () {
    const [count, setCount] = useState(0);

    const increment = () => {
        setCount(count+1)
    }
    const decrement = () => {
        setCount(count-1)
    }
    const reset = () =>{
        setCount(0)
    }

    return (
        <div className="counter-container">
            <p className="counter-p">{count}</p>
            <button className="counter-b" onClick={decrement}>Decrement</button>
            <button className="counter-b" onClick={reset}>Reset</button>
            <button className="counter-b" onClick={increment}>Increment</button>
            {/* <input onChange={reset}>On Change</input> */}
            {/* <p>Text</p> */}
        </div>
    )

}

export default Counter