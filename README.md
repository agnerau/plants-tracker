**GOplants**

user can document the growth of his plants

endpoints:
CRUD plant
create, get height (get height - all heights)

CRUD diseases, fertilizers - nice to have

frontend - nice to have
show each plant with a photo or default image in a card
when pressing on a card redirects to plant page with info and height graph

**Tables**

PLANTS

- bought_at - null
- planted_at - null
- name
- nickname - null
- died_at - null

HEIGHTS
- value
- date
- fk_plant

DISEASES
- name
- started_at
- ended_at - null
- fk_plant

FERTILIZERS
- name
- started_at
- ended_at - null
- fk_plant

