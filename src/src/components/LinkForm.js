import React from "react";

function LinkForm({linkValue, setLinkValue}) {
    const handleChange = (e) => {
        if (e !== null && e !== undefined) {
            const value = e.target.value;
            setLinkValue(value);
        }
    }

    return (
        <input type="text" placeholder="Insert your starting link here" className="inoneline" onChange={handleChange} style={{ margin: '10px',width:'500px',height:'30px',borderRadius:'20px',padding:'5px' }}required />
    );
};

export default LinkForm;