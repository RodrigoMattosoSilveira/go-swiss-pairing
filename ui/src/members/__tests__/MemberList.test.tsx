import { render, screen } from '@testing-library/react';
import React from 'react';
import { MemoryRouter } from 'react-router-dom';
import MemberList from '../MemberList';
import { MOCK_MEMBERS } from '../MockMembers';
import exp from "constants";
import userEvent from '@testing-library/user-event';
// import { Provider } from 'react-redux';
// import { store } from '../../state';

describe('<MemberList />', () => {
    let onSave: jest.Mock;
    const setup = () => {
        render(
            // <Provider store={store}>
            <MemoryRouter>
                <MemberList members={MOCK_MEMBERS} onSave={onSave}/>
            </MemoryRouter>
            // </Provider>
        );

    }
    beforeEach(() => {});

    test('should render without crashing', () => {
        setup();

        expect(screen).toBeDefined();
    });
    test('should display list', () => {
        setup();
        expect(screen.getAllByRole('button')).toHaveLength(MOCK_MEMBERS.length);
        expect(screen.getAllByRole('generic', { name: 'member.first' })).toHaveLength(MOCK_MEMBERS.length);
        expect(screen.getAllByRole('generic', { name: 'member.last' })).toHaveLength(MOCK_MEMBERS.length);
        expect(screen.getAllByRole('generic', { name: 'member.email' })).toHaveLength(MOCK_MEMBERS.length);
        expect(screen.getAllByRole('generic', { name: 'member.cell' })).toHaveLength(MOCK_MEMBERS.length);
        expect(screen.getAllByRole('generic', { name: 'member.rating' })).toHaveLength(MOCK_MEMBERS.length);
    });
    test('should display form when edit clicked', async () => {
       setup();
       // eslint-disable-next-line testing-library/render-result-naming-convention
           const user = userEvent.setup();
       await user.click(
             screen.getByRole('button', {name:"m5lpUgrpH"})
           );
       expect(
             screen.getByRole('form', {
                   name: /edit a member/i,
             })
       ).toBeInTheDocument();
    });
})