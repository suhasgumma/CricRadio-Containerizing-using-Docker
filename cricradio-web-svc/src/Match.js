import React from "react";
import { Link } from "react-router-dom";

const Match = (props) =>{
    const {teams,details,matchId,seriesId,onMatchClick} = props;
    return (
        <div className="sidebar-contact" onClick={()=>{onMatchClick({"teams":teams,"details":details,"matchId":matchId,})}}>
            <div>
                <div>
                    <span style={{fontFamily:"Poppins",fontWeight:"500",color:"#E9EDEF",fontSize:"14px"}}>{teams}</span>
                </div>
                <div>
                    <span style={{fontFamily:"Poppins",fontWeight:"400",color:"#D1D7DB",fontSize:"9px"}}>{details}</span>
                </div>

            </div>
        </div>
    )
}

export default Match