import React, { useState } from 'react';
import { signOut } from "firebase/auth";
import { getAuth, createUserWithEmailAndPassword } from "firebase/auth";
import { signInWithEmailAndPassword } from "firebase/auth";
import { fireAuth } from "./firebase";

const LoginForm: React.FC = () => {
  const [userName, setUserName] = useState<string | null>(null);
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");

  const createuser = () => {
    const auth = getAuth();
    createUserWithEmailAndPassword(auth, email, password)
    .then((userCredential) => {
    const user = userCredential.user;
    alert("ログインユーザー: " + user.displayName);
    setUserName(user.displayName)
    })
    .catch((error) => {
    const errorMessage = error.message;
    alert(errorMessage);
  });
  }

  const signInWithEmailandPassword = () => {
    const auth = getAuth();
    signInWithEmailAndPassword(auth, email, password)
    .then((userCredential: { user: any; }) => {
      // Signed in 
      const user = userCredential.user;
      alert("ログインユーザー: " + user.displayName);
      setUserName(user.displayName)
      }).catch((err: any) => {
        alert(err);
  });
  }

  const signOutWithEmailAndPassword = () => {
    signOut(fireAuth).then(() => {
      alert("ログアウトしました");
    }).catch(err => {
      alert(err);
    });
  }
  const handleSubmit = async (e: { preventDefault: () => void; }) => {
    e.preventDefault()
  }

  return (
    <div>
      <button onClick={signInWithEmailandPassword}>
        メール・パスワードでログイン
      </button>
      <button onClick={signOutWithEmailAndPassword}>
        ログアウト
      </button>
      <form style={{ display: "flex", flexDirection: "column" }} onSubmit={handleSubmit}>
        <label>Email: </label>
        <input
          type="string"
          value={email}
          onChange={(e) => setEmail(e.target.value)}
        ></input>
      </form>
      <form style={{ display: "flex", flexDirection: "column" }} onSubmit={handleSubmit}>
        <label>Password: </label>
        <input
          type="string"
          value={password}
          onChange={(e) => setPassword(e.target.value)}
        ></input>
        <button type={"submit"} onClick={createuser} >新規ユーザー作成</button>
      </form>
    </div>
    
  );
};
export default LoginForm;