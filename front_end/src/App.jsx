import FetchTest from "./FetchTest";
import Background from "./Background/background.jsx";
import FriendsList from "./Buddies/FriendsList";

function App() {
    return (
        <>  
        <FriendsList></FriendsList>
            <Background></Background>
            <h1 className="text-3xl font-bold underline">
                Hey guys! This was styled with Tailwind CSS!
                <br /> <br /> - Omar
            </h1>
            {/* <FetchTest/> */}
        </>
    );
}

export default App;
