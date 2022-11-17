import React from 'react';
import { Member } from './Member';
import MemberCard from "./MemberCard";
import MemberForm from "./MemberForm";

interface ProjectListProps {
    members: Member[];
}

function ProjectList({ members }: ProjectListProps) {
    // return <pre>{JSON.stringify(members, null, ' ')}</pre>;
   return (
        <div className="row">
            {members.map((member) => (
            <div key={member.id} className="cols-sm">
                <MemberCard member={member}></MemberCard>
                <MemberForm/>
            </div>
            ))}
        </div>
   );
}

export default ProjectList;