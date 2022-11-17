import React from 'react';
import { Member } from './Member';

interface ProjectListProps {
    members: Member[];
}

function ProjectList({ members }: ProjectListProps) {
    return <pre>{JSON.stringify(members, null, ' ')}</pre>;
}

export default ProjectList;