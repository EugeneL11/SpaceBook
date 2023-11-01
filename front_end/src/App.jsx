import FetchTest from "./FetchTest.jsx";
import Background from "./Background/background.jsx";

function App() {
    return (
        <>
            <Background></Background>
            <h1 className="text-3xl font-bold underline">
                Hey guys! This was styled with Tailwind CSS!
                <br /> <br /> - Omar
            </h1>
            <FetchTest />
        </>
    );
}

export default App;
