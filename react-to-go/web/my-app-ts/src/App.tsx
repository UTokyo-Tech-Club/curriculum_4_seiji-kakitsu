import React from "react";
import "./App.css";
import { useState } from "react";
import { useEffect } from "react";

interface UserData {
  name: string;
  age: number;
}

function App() {
  const [age, setAge] = useState<number>(0);
  const [name, setName] = useState<string>("");
  const [userData, setUserData] = useState<UserData[]>([]);

  const handleSubmit = async (e: { preventDefault: () => void; }) => {
    e.preventDefault()
    if (!name) {
      alert("Please enter name");
      return;
    }

    if (name.length > 50) {
      alert("Please enter a name shorter than 50 characters");
      return;
    }

    if (age < 20 || age > 80) {
      alert("Please enter age between 20 and 80");
      return;
    }

    try{
      const response = await fetch(
        "http://localhost:8000/user",
        {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify({
            name,
            age,
          }),
        }
      );
      if (response.status === 200) {
        fetchUsers();
      } else {
        console.error("POST request failed")
      }

      setName("");
      setAge(0);
    } catch (err) {
      console.error(err)
    }
  };
  const fetchUsers = async () => {
    try{
      const getResponse = await fetch("http://localhost:8000/user", {
        method: "GET",
        headers: {
          "Content-Type": "application/json",
        },
      });
    
      if (getResponse.status === 200) {
        // GETリクエストの結果を処理
        const userData = await getResponse.json();
        setUserData(userData);
        // userDataを適切に処理するコードをここに追加
      } else {
        // GETリクエストが失敗した場合の処理
        console.error("GET request failed");
      }
    } catch (err) {
      console.error(err)
    }
  };

  useEffect(() => {
    fetchUsers();
  },[]);

  return (
    <div className="App">
      <header className="App-header">
        <div className="border">
          <p className="p">User Register </p>
        </div>
        <p>
          Edit <code>src/App.js</code> and save to reload.
        </p>
        <a
          className="App-link"
          href="https://reactjs.org"
          target="_blank"
          rel="noopener noreferrer"
        >
        </a>
      </header>
        <form style={{ display: "flex", flexDirection: "column" }} onSubmit={handleSubmit}>
          <div className="block_1">
            <label>Name: <input
              type={"text"}
              value={name}
              onChange={(e) => setName(e.target.value)}
              className="txt_1"
            ></input></label>
          </div>
          <div className="block_2">
          <label>Age: <input 
            type="number"
            value={age}
            onChange={(e) => setAge(parseInt(e.target.value))}
            className="txt_1"
          ></input></label>
          </div>
          <button type={"submit"} className="button">POST</button>
        </form>
        <div>
          {userData.map((user, index) => (
            <div key={index} className="txt_2">
              <p>{user.name}, {user.age}</p>
            </div>
          ))}
        </div>
    </div>
    );
  }

  export default App;
