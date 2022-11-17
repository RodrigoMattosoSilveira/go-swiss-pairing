import { Member } from './Member';
import React from 'react';

interface ProjectCardProps {
    member: Member;
}

function MemberCard(props: ProjectCardProps) {
    const { member: member } = props;
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
            </section>
        </div>
    );
}

export default MemberCard;