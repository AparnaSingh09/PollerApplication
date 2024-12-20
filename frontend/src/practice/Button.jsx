function Button() {
     let count = 0;
    const handleClick = (name) => {
        if (count < 3){
count++;
console.log(`${name} you clicked me ${count} times`)
        }
        console.log("OUCH!!")
    }
    return (<button onClick={() => handleClick("girl")}>Click me</button>);
}

export default Button