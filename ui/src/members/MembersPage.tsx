import React from 'react';
import {MOCK_MEMBERS} from "./MOCK_MEMBERS";
import MemberList from "./MemberList";

function ProjectsPage() {
    return (
        <>
            <h1>Projects</h1>
            {/*<pre>{JSON.stringify(MOCK_MEMBERS, null, ' ')}</pre>*/}
            <MemberList members={MOCK_MEMBERS} />
        </>
        );
}

export default ProjectsPage;