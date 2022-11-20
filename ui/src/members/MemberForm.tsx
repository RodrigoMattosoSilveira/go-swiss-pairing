import React, { SyntheticEvent, useState } from 'react';
import { Member } from "./Member"
interface MemberFormProps {
    member: Member;
    onSave: (member: Member) => void;
    onCancel: () => void;
}

function MemberForm({member: initialMember, onSave, onCancel}: MemberFormProps) {
    // const {member: initialMember, onSave, onCancel} = props;
    const [member, setMember] = useState(initialMember);
    const [errors, setErrors] = useState({
        first: '',
        last: '',
        email: '',
        cell: '',
        rating: ''
    });
    const handleSubmit = (event: SyntheticEvent) => {
        console.log("handling submit")
        event.preventDefault();
        if (!isValid()) return;
        onSave(member);
    };
    const handleChange = (event: any) => {
        const { type, name, value, checked } = event.target;
        // if input type is checkbox use checked
        // otherwise it's type is text, number etc. so use value
        let updatedValue = type === 'checkbox' ? checked : value;

        //if input type is number convert the updatedValue string to a +number
        if (type === 'number') {
            updatedValue = Number(updatedValue);
        }
        const change = {
            [name]: updatedValue,
        };

        let updatedMember: Member;
        // need to do functional update b/c
        // the new project state is based on the previous project state
        // so we can keep the project properties that aren't being edited +like project.id
        // the spread operator (...) is used to
        // spread the previous project properties and the new change
        setMember((p) => {
            updatedMember = new Member({ ...p, ...change });
            return updatedMember;
        });
        setErrors(() => validate(updatedMember));
    };
    function validate(member: Member) {
        let errors: any = { first: '', last: '', email: '', cell: '', rating: '' };
        if (member.first.length === 0) {
            errors.first = 'First name is required';
        }
        if (member.first.length > 0 && member.first.length < 3) {
            errors.first = 'First name needs to be at least 3 characters.';
        }

        if (member.last.length === 0) {
            errors.last = 'Last last is required';
        }
        if (member.last.length > 0 && member.last.length < 3) {
            errors.last = 'Last name needs to be at least 3 characters.';
        }

        if (member.email.length === 0) {
            errors.email = 'Email is required';
        }

        if (member.cell.length === 0) {
            errors.cell = 'Cell number is required';
        }

        if (member.rating === 0 || member.rating > 3000) {
            errors.rating = 'Member rating must be between 0 and 3000';
        }

        return errors;
    }
    function isValid() {
      return (
        errors.first.length === 0 &&
        errors.last.length === 0 &&
        errors.email.length === 0 &&
        errors.cell.length === 0 &&
        errors.rating.length === 0
      );
    }
    return (
        <form className="input-group vertical"
              onSubmit={handleSubmit}
        >
            <label htmlFor="first">First Name</label>
            <input
                type="text"
                name="first"
                placeholder="enter first name"
                value={member.first}
                onChange={handleChange}
            />
            {errors.first.length > 0 && (
                <div className="card error">
                    <p>{errors.first}</p>
                </div>
           )}

            <label htmlFor="last">First Name</label>
            <input
                type="text"
                name="last"
                placeholder="enter last name"
                value={member.last}
                onChange={handleChange}
            />
            {errors.last.length > 0 && (
                <div className="card error">
                    <p>{errors.last}</p>
                </div>
            )}

            <label htmlFor="email">Email</label>
            <input
                type="text"
                name="email"
                placeholder="enter email"
                value={member.email}
                onChange={handleChange}
            />
            {errors.email.length > 0 && (
                <div className="card error">
                    <p>{errors.email}</p>
                </div>
            )}

            <label htmlFor="cell">Cell</label>
            <input
                type="text"
                name="cell"
                placeholder="enter cell number"
                value={member.cell}
                onChange={handleChange}
            />
            {errors.cell.length > 0 && (
                <div className="card error">
                    <p>{errors.cell}</p>
                </div>
            )}

            <label htmlFor="rating">Rating</label>
            <input
                type="number"
                name="rating"
                placeholder="enter rating"
                value={member.rating}
                onChange={handleChange}
            />
            {errors.rating.length > 0 && (
                <div className="card error">
                    <p>{errors.rating}</p>
                </div>
            )}

            <label htmlFor="isActive">Active</label>
            <input
                type="checkbox"
                name="isActive"
                checked={member.isActive}
                onChange={handleChange}
            />

            <div className="input-group">
                <button className="primary bordered medium">Save</button>
                <span />
                <button
                    type="button"
                    className="bordered medium"
                    onClick={onCancel}
                >
                    cancel
                </button>
            </div>
        </form>
    );
}

export default MemberForm;