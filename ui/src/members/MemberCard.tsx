import { Member } from './Member';
import React from 'react';

interface ProjectCardProps {
    member: Member;
}

function MemberCard(props: ProjectCardProps) {
    const { member: member } = props;
    const handleEditClick = (memberBeingEdited: Member) => {
          console.log(memberBeingEdited);
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