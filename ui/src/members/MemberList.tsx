import React from 'react';
import { Member } from './Member';
import MemberCard from "./MemberCard";
import MemberForm from "./MemberForm";

interface ProjectListProps {
    members: Member[];
}

function ProjectList({ members }: ProjectListProps) {
    const handleEdit = (member: Member) => {
        console.log(member);
    }
    // return <pre>{JSON.stringify(members, null, ' ')}</pre>;
   return (
        <div className="row">
            {members.map((member) => (
            <div key={member.id} className="cols-sm">
                <MemberCard member={member} onEdit={handleEdit}></MemberCard>
                <MemberForm/>
            </div>
            ))}
        </div>
   );
}

export default ProjectList;