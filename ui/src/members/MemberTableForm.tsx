import React from "react";

function MemberTableForm() {
    return (
        <form className="input-group vertical">
            <label htmlFor="name">Member First Name</label>
            <input type="text" name="first" placeholder="enter first name" />

            <label htmlFor="description">Member Last Name</label>
            <input type="text" name="last" placeholder="enter last name" />

            <label htmlFor="description">Member Email</label>
            <input type="text" name="email" placeholder="enter email" />

            <label htmlFor="description">Member Cell</label>
            <input type="text" name="cell" placeholder="enter cell" />

            <label htmlFor="isActive">Active?</label>
            <input type="checkbox" name="isActive" />
            <div className="input-group">
                <button className="primary bordered medium">Save</button>
                <span />
                <button type="button" className="bordered medium">
                    cancel
                </button>
            </div>
        </form>
    );
}

export default MemberTableForm;