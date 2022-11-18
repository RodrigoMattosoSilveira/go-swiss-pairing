import { Member } from './Member';
import React from 'react';

interface MemberCardProps {
    member: Member;
    onEdit: (member: Member) => void
}

function MemberCard(props: MemberCardProps) {
    const { member: member, onEdit } = props;
    const handleEditClick = (memberBeingEdited: Member) => {
        // console.log(memberBeingEdited);
        onEdit(memberBeingEdited)
    };
    return (
        <div className="card">
            <img src={"assets/" + member.imageUrl} alt={member.first + " " + member.last} />
            <section className="section dark">
                <h5 className="strong">
                    <strong>{member.first + " " + member.last}</strong>
                </h5>
                <p>Cell: {member.cell}</p>
                <p>Email {member.email}</p>
                <p>Rating : {member.rating.toLocaleString()}</p>
                <button
                    className="bordered"
                    onClick={() => handleEditClick((member))}
                >
                    <span className="icon-edit "></span>
                    Edit
                </button>
            </section>
        </div>
    );
}

export default MemberCard;