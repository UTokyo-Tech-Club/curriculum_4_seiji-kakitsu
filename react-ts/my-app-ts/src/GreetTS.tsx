import React from "react";
export const GreetTS = () => {
    const greet = (name: string) => {
    return "Hello, " + name + "!!";
  };
  
  return (
    <div>
      <p>{greet("John")}</p>
      <p>{greet("Taro")}</p>　//引数に数字（number型）が入っているためエラーが出る
    </div>
  )
}