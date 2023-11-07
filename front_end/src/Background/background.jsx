import {
    Sparkles,
    Stars,
  } from "@react-three/drei";
import { Canvas } from "@react-three/fiber";

export default function Background() {
    const colorArray = [
        "#fffa86",
        "#000000",
    ];
  
    const count = 300;

    return (
        <div className="h-screen w-screen fixed !-z-20">
        <Canvas className="!-z-20">
    
            <Sparkles
                /** Number of particles (default: 100) */
                count={count}
                /** Speed of particles (default: 1) */
                speed={1}
                /** Opacity of particles (default: 1) */
                // opacity?: number | Float32Array
                /** Color of particles (default: 100) */
                color={ "#fffa86" }
                /** Size of particles (default: randomized between 0 and 1) */
                size={3}
                /** The space the particles occupy (default: 1) */
                scale={10}
                /** Movement factor (default: 1) */
                // noise?: number | [number, number, number] | THREE.Vector3 | Float32Array
                className="!-z-20"
            />

            <Stars radius={1} depth={50} count={1000} factor={5} saturation={0} fade speed={5} className="!-z-20"/>

        </Canvas>  
        </div>
    );
}