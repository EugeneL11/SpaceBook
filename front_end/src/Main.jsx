import "./index.css";
import App from "./App";
import { Canvas } from "@react-three/fiber";
import Background from "./Background/background";

export default function Main() {
  return (
    <>
      <div className="h-screen w-screen fixed">

        <Canvas
        className="!absolute three"
        >
          <Background/>
        </Canvas>  
      </div>
      
      <App  />
    </>
  );
}