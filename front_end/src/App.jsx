import FetchTest from "./FetchTest";
import Background from "./Background/background.jsx";
import MyProfile from "./Profile/MyProfile";

function App() {
    return (
        <>  
            <Background></Background>
            <MyProfile></MyProfile>
            <h1 className="text-3xl font-bold underline">
                Hey guys! This was styled with Tailwind CSS!
                <br /> <br /> - Omar
            </h1>
            {/* <FetchTest/> */}
        </>
    );
}

export default App;
