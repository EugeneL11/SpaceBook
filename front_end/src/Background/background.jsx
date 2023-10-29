import { useRef } from "react";
import {
    //   useHelper,
    Sparkles,
    Stars,
  } from "@react-three/drei";
import { useThree, Canvas } from "@react-three/fiber";

export default function Background() {
    const colorArray = [
        "#fffa86",
        "#000000",
    ];
  
    const count = 500;

    return (
        <div className="h-screen w-screen fixed">

        <Canvas>
    
          <Sparkles
            /** Number of particles (default: 100) */
            count={count}
            /** Speed of particles (default: 1) */
            speed={0.5}
            /** Opacity of particles (default: 1) */
            // opacity?: number | Float32Array
            /** Color of particles (default: 100) */
            // random color
            color={ "#fffa86" }
            /** Size of particles (default: randomized between 0 and 1) */
            size={3}
            /** The space the particles occupy (default: 1) */
            scale={10}
            /** Movement factor (default: 1) */
            // noise?: number | [number, number, number] | THREE.Vector3 | Float32Array
          />

        <Stars radius={1} depth={50} count={5000} factor={4} saturation={0} fade speed={5} />


        </Canvas>  
        </div>

      );
    }