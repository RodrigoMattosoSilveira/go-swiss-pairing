import React from 'react';
import { render, screen } from '@testing-library/react';
import { MemoryRouter } from 'react-router-dom';
import { Member } from '../Member';
import MemberForm from '../MemberForm';
import userEvent from '@testing-library/user-event';

describe('<MemberForm />', () => {
    let member: Member;
    let updatedMember: Member;
    let handleCancel = jest.fn();
    let firstTextBox: any;
    let lastTextBox: HTMLElement;
    let emailTextBox: HTMLElement;
    let cellTextBox: HTMLElement;
    let ratingTextBox: HTMLElement;
    let isActiveCheckbox: HTMLElement;
    let onSave = jest.fn();

    const setup = () => {
        render(
            <MemoryRouter>
                <MemberForm member={member} onSave={onSave} onCancel={handleCancel} />
            </MemoryRouter>
        );

        firstTextBox = screen.getByRole('textbox', {
            name: /member first name/i,
        });
        lastTextBox = screen.getByRole('textbox', {
            name: /member last name/i,
        });
        emailTextBox = screen.getByRole('textbox', {
            name: /member email address/i,
        });
        cellTextBox = screen.getByRole('textbox', {
            name: /member cell phone number/i,
        });
        ratingTextBox = screen.getByRole('spinbutton', {
            name: /member rating/i,
        });
        isActiveCheckbox = screen.getByRole('checkbox', {
            name: /member status/i,
        });
    };

    beforeEach(() => {
        member = new Member({
            id: "m5lpUgrpH",
            first: 'Adeline',
            last: 'Hodge',
            email: "Adeline.Hodge@yahoo.com",
            password: "oPT14J30I#y4",
            cell: "801 277-6891",
            rating: 2043,
            isActive: true
        });
        updatedMember = new Member({
            id: "m5lpUgrpH",
            first: 'Francine',
            last: 'Hodge',
            email: "Adeline.Hodge@yahoo.com",
            password: "oPT14J30I#y4",
            cell: "209 290-5244",
            rating: 1658,
            isActive: true
        });
        handleCancel = jest.fn();
    });

    test('should load member into form', () => {
        setup();
        expect(
            screen.getByRole('form', {
                name: /Edit a Member/i,
            })
        ).toHaveFormValues({
            first: member.first,
            last: member.last,
            email: member.email,
            cell: member.cell,
            rating: member.rating,
            isActive: member.isActive
        });
    });
    test('should accept input', async () => {
        setup();
        const user = userEvent.setup();
        await user.clear(firstTextBox);
        await user.type(firstTextBox, updatedMember.first);
        expect(firstTextBox).toHaveValue(updatedMember.first);

        await user.clear(cellTextBox);
        await user.type(cellTextBox, updatedMember.cell);
        expect(cellTextBox).toHaveValue(updatedMember.cell);

        // In order fot this to work, we have to overcome the browser's autocomplete behavior. The suggestions to use
        // autoComplete="false", or variations thereof do not work.
        // Playing with the actual form, I observed that I can select the rating value by double-clicking on any
        // position to the left of any the digits and then, with the rating highlighted I was able to type any value.
        // Below is the logic that implements this behavior.
        user.pointer({target: ratingTextBox, offset: 1})
        user.dblClick(ratingTextBox)
        await user.type(ratingTextBox, updatedMember.rating.toString());
        expect(ratingTextBox).toHaveValue(updatedMember.rating);
    });
    test('first should display required validation', async () => {
          setup();
          const user = userEvent.setup();
          await user.clear(firstTextBox);
          expect(screen.getByRole('alert')).toBeInTheDocument();
    });
    test('first should display minlength validation', async () => {
          setup();
          const user = userEvent.setup();
          await user.clear(firstTextBox);
          await user.type(firstTextBox, 'ab');
          await expect(screen.getByRole('alert')).toBeInTheDocument();
          await user.type(firstTextBox, 'c');
          expect(screen.queryByRole('alert')).not.toBeInTheDocument();
    });
    // Note that I had to implement an interation to highlight the rating
    test('budget should display not 0 validation', async () => {
        setup();
        const user = userEvent.setup();
        await user.clear(ratingTextBox);

        // Test for an invalid rating
        user.pointer({target: ratingTextBox, offset: 1})
        user.dblClick(ratingTextBox)
        await user.type(ratingTextBox, '1199');
        expect(screen.getByRole('alert')).toBeInTheDocument();

        // Test for a valid rating
        user.pointer({target: ratingTextBox, offset: 1})
        user.dblClick(ratingTextBox)
        await user.type(ratingTextBox, '1200');
        expect(screen.queryByRole('alert')).not.toBeInTheDocument();
    });
});