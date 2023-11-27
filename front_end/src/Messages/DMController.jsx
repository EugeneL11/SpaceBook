import { React, useState, useEffect} from "react";
import DMList from "./DMList";
import DMMessage from "./DMMessage";
import NewDM from "./NewDM"
import OtherProfile from "../Profile/OtherProfile";

function DMController(props) {
    const toggleHomepage = props.toggleHomepage
    const [DMstate, setDMstate] = useState(null);
    useEffect(() => {
        if (props.wormhole) {
            toggleDMMessage(props.friendID)
        } else {
            toggleDMList()
        }
    }, [])

    function toggleDMList(){
        console.log(toggleHomepage)
        setDMstate(<DMList toggleHomepage={toggleHomepage} toggleDMMessage={toggleDMMessage} toggleNewDM = {toggleNewDM}/>)
    }
    function toggleDMMessage(friendID){
        console.log("Whastup")
        setDMstate(<DMMessage friendID={friendID} toggleDMList={toggleDMList} toggleOtherProfile={toggleOtherProfile} toggleDMMessage={toggleDMMessage}/>)
    }
    function toggleOtherProfile(userID,backEvent){
        setDMstate(<OtherProfile userID = {userID} goBackScreen = {backEvent}/>)
    }
    // make it toggle other profile properly
    function toggleNewDM(friendID){
        setDMstate(<NewDM friendID={friendID} toggleDMMessage={toggleDMMessage} toggleDMList = {toggleDMList} toggleOtherProfile = {toggleOtherProfile}/>)
    }
    return (<div>{DMstate}{/* */}</div>);

}

export default DMController;