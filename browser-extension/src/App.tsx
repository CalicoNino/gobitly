import React, { useState } from "react";
import logo from "./logo.png";

const App = () => {
  const [copied, setCopied] = useState(false);
  const [gobitly, setGobitly] = useState("");
  const [userInput, setUserInput] = useState("");

  const currentUrl = window.location.href.split("//")[1];
  const inputValue = userInput === "" ? currentUrl : userInput;

  const copy = () => {
    navigator.clipboard.writeText(gobitly);
    setCopied(true);
  };

  const generate = () => {
    setGobitly(inputValue);
  };

  return (
    <div className="w-48 h-full py-3 flex flex-col gap-5 justify-center items-center bg-blue-300 border border-blue-800 inset-10">
      <img src={logo} className="w-20 h-20" alt="logo" />
      <div className="flex flex-col p-3 gap-5 w-36 justify-center items-center">
        <div className="relative h-11 w-full">
          <input
            placeholder="Your Url"
            onChange={(value) => {
              setUserInput(value.target.value);
            }}
            value={inputValue}
            className="peer h-full w-full border-b border-blue-gray-200 bg-transparent pt-4 pb-1.5 font-sans text-sm font-bold text-blue-900 outline outline-0 transition-all placeholder-shown:border-blue-gray-200 focus:border-blue-500 focus:outline-0 disabled:border-0 disabled:bg-blue-gray-50"
          />
          <label className="after:content[' '] pointer-events-none absolute left-0 -top-2.5 flex h-full w-full select-none text-sm font-normal leading-tight text-blue-900 transition-all after:absolute after:-bottom-2.5 after:block after:w-full after:scale-x-0 after:border-b-2 after:border-blue-500 after:transition-transform after:duration-300 peer-placeholder-shown:leading-tight peer-placeholder-shown:text-blue-gray-500 peer-focus:text-sm peer-focus:leading-tight peer-focus:text-blue-500 peer-focus:after:scale-x-100 peer-focus:after:border-blue-500 peer-disabled:text-transparent peer-disabled:peer-placeholder-shown:text-blue-500">
            URL
          </label>
        </div>

        <div className="w-32 flex flex-col justify-center items-center overflow-x-auto overflow-y-hidden bg-transparent text-center text-blue-900 font-extrabold">
          <p className="w-full px-1 mx-1">
            {gobitly === "" ? "GoBitly" : gobitly}
          </p>
        </div>
        <div className="flex flex-row w-full gap-2 justify-center items-center">
          <button
            onClick={generate}
            className={`w-24 h-8 bg-transparent border border-blue-900 hover:bg-blue-500 rounded-md ${
              copied && "bg-blue-400"
            }`}
          >
            Generate
          </button>
          <button
            onClick={copy}
            className={`w-8 h-8 bg-transparent border border-blue-900 hover:bg-blue-500 rounded-md ${
              copied && "bg-blue-400"
            }`}
          >
            {copied ? "✔️" : "📋"}
          </button>
        </div>
      </div>
    </div>
  );
};

export default App;
