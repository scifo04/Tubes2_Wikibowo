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
        <button onClick={handleClick}>Start Racing</button>
    );
}

export default SearchButton;