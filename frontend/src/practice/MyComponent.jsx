import React,{useState} from "react";

function MyComponent() {
    
    let [name, setName] = useState("old");

    const updateName  = () => {
        setName("appyy");
    }

    return(<div>
        <p>Name: {name}</p>
        <button onClick={updateName}>Set name</button>
    </div>)

}


export default MyComponent