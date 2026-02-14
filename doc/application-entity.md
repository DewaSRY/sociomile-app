# Database entity

user_role
- name

user
- name
- email
- password
- created_at
- updated_at
- organization_id
- role_id

organization
- name
- owner_id

conversation
- organization_id
- guess_id
- organization_staff_id
- created_at
- status  // pending , in_progress, done

conversation_message
- organization_id
- conversation_id
- created_by_id
- created_at
- message

ticket
- organization_id
- conversation_id
- created_by_id
- created_at
- ticket_number
- name
- status  // pending , in_progress, done