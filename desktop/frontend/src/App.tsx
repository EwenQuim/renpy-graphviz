import { useState } from "react";
import logo from "./assets/images/logo-universal.png";
import "./App.css";
import { Greet } from "../wailsjs/go/main/App";
import { Graphviz } from "graphviz-react";

const App = () => {
  const [resultText, setResultText] = useState("Renpy Graphviz");
  const [name, setName] = useState("");

  const updateName = (e: any) => setName(e.target.value);

  const greet = async () => {
    const newName = await Greet(name);
    setResultText(newName);
  };

  return (
    <div
      id="App"
      style={{
        display: "flex",
        alignItems: "center",
      }}
    >
      <aside
        id="input"
        style={{
          backgroundColor: "beige",
          color: "black",
          padding: "1rem",
          top: "0",
          left: "0",
          height: "100vh",
          borderRight: "1px solid black",
        }}
      >
        <h1 id="result" className="result">
          Renpy Graphviz
        </h1>
        <input
          id="name"
          className="input"
          onChange={updateName}
          autoComplete="off"
          name="input"
          type="text"
        />
        <button className="btn" onClick={greet}>
          Greet
        </button>
      </aside>
      <Graphviz
        dot={toDisplay}
        options={{
          width: "100%",
          zoom: true,
          engine: "dot",
        }}
      />
    </div>
  );
};

export default App;

const toDisplay = `digraph  {

    n2[color="red",label="bad ending",shape="septagon",style="bold"];
    n4[color="purple",fontsize="16",label="GOOD ENDING",shape="rectangle",style="bold"];
    n5[label="route 2"];
    n3[label="routeAlternative"];
    n1[color="purple",fontsize="16",label="ROUTEONE",shape="rectangle",style="bold"];
    n6[color="purple",fontsize="16",label="STAAA AA 6 RT",shape="rectangle",style="bold"];
    n5->n2[label="",style="dotted"];
    n3->n4[label=""];
    n1->n2[label=""];
    n1->n3[label="",style="dotted"];
    n6->n1[label="Ah! D'accord, si vous le dites..."];
    n6->n5[label="Vous êtes sûr?"];
    n6->n3[label="More"];

}`;
