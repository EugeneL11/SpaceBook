import { React, useState } from "react";
import DMList from "./DMList";
import DMMessage from "./DMMessage";
import OtherProfile from "../Profile/OtherProfile";

function DMController(props) {
    const toggleHomepage = props.toggleHomepage
    const [DMstate, setDMstate] = useState(<DMList toggleHomepage={toggleHomepage} toggleDMMessage={toggleDMMessage}/>);
    function toggleDMList(){
        setDMstate(<DMList toggleHomepage={toggleHomepage} toggleDMMessage={toggleDMMessage}/>)
    }
    function toggleDMMessage(friendID){
        console.log("Whastup")
        setDMstate(<DMMessage friendID={friendID} toggleDMList={toggleDMList} toggleOtherProfile={toggleOtherProfile}/>)
    }
    function toggleOtherProfile(userID,backEvent){
        setDMstate(<OtherProfile userID = {userID} goBackScreen = {backEvent}/>)
    }
    return (<div>{DMstate}{/* toggleDMMessage={toggleDMMessage}*/}</div>);

}

export default DMController;