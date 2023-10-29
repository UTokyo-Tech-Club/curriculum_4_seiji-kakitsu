import React from "react";
import logo from "./logo.svg";
import "./App.css";
import { useState } from "react";

function App() {
  const [age, setAge] = useState(0);

  const handleSubmit = async (e: { preventDefault: () => void; }) => {
    e.preventDefault()
    const response = await fetch(
      "https://test2-9881e-default-rtdb.firebaseio.com/.json",
      {
        method: "GET",
        headers: {
          "Content-Type": "application/json",
        },
      }
    );
    interface UserData {
      name: string;
      age: number;
    }
    const data:UserData = await response.json();
    const obj = Object.values(data).filter((v) => v.name === "inada")[0];
    if (typeof obj.age === "number") {
      setAge(obj.age + 10);
    } else {
      setAge(parseInt(obj.age) + 10);
    }
  };

  return (
    <div className="App">
      <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
        <p>
          Edit <code>src/App.js</code> and save to reload.
        </p>
        <a
          className="App-link"
          href="https://reactjs.org"
          target="_blank"
          rel="noopener noreferrer"
        >
          Learn React
        </a>
      </header>
      <form style={{ display: "flex", flexDirection: "column" }} onSubmit={handleSubmit}>
            <label>Age: </label>
            <input
              type="number"
              value={age}
              onChange={(e) => setAge(parseInt(e.target.value))}
            ></input>
            <button type={"submit"}>Submit</button>
          </form>
      </div>
    );
  }

  export default App;