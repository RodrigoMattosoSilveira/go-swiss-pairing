import React from 'react';
import { Member } from './Member';

interface MemberListProps {
    members: Member[];
}

function ProjectList({ members }: MemberListProps) {
    return <pre>{JSON.stringify(members, null, ' ')}</pre>;
}

export default ProjectList;