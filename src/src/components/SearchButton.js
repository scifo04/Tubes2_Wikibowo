import React from "react";

function SearchButton({isOn,linkValue,finValue,tempLinks,setTempLinks}) {
    const handleClick = async () => {
        try {
            const response = await fetch('http://localhost:8000', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({
                    linkValue: linkValue,
                    finValue: finValue,
                    isOn: isOn,
                }),
            });

            if (!response.ok) {
                throw new Error('Network response was not ok');
            }

            const tempResponse = await response.json();
            setTempLinks(tempResponse.links);
        } catch (error) {
            console.error('Error:', error);
        }
    };

    return (
        <button onClick={handleClick} style={{fontFamily:"Poppins",fontSize:"20px",border:"2px solid white",borderRadius:"10px",backgroundColor:"black",color:"white",width:"150px",height:"40px"}}>Start Racing</button>
    );
}

export default SearchButton;