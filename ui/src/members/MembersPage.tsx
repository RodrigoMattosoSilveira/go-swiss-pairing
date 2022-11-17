import React from 'react';
import {MOCK_MEMBERS} from "./MOCK_MEMBERS";

function ProjectsPage() {
    return (
         <>
               <h1>Projects</h1>
              <pre>{JSON.stringify(MOCK_MEMBERS, null, ' ')}</pre>
             </>
        );
}

export default ProjectsPage;