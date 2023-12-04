# Technical test - Back end

# Context & Hypothesis

At Side, the operational recruitment team (or ops team) is in charge of proposing the right temp workers for the right task. In order to do so, this team is using different internal tools that we built over time.

During this test, let’s act as if you were building something for this team to improve its workflow.

### Side’s current global workflow

1. **Companies post tasks on side.co** - A company is asking for X temp workers, working at a specific location, from a StartDate to an EndDate on a TypeOfTask task.
2. **Ops team is reviewing the tasks** - In order to make sure that everything is properly formatted and push tasks to Side’s workers base.
3. **Tasks are visible to workers** - Through their Side mobile app.
4. **Temp workers are applying for tasks** - Depending on their availability, willingness to work on this task, task pricing, task location and so on...
5. **Application time is over** - Once enough (qualified) workers did apply for the task, the task is removed from visibility.
6. **The “selection” process begins** - ****Ops are reviewing candidates and selecting the best ones so that companies get the best profiles.
7. **Selection is done** - Selected workers are notified that they will be working. Unselected workers are also notified and can apply for something else. Companies are notified that workers matching their criteria were found and can access their contact information.
8. **Work**
9. **Invoicing/Payroll**
10. **Payment**

*Keep in mind that this workflow (from 1. to 7.)  happens many times a day and that each task evolves in a different time frame.*

*As an example, a task for 1 worker that’ll start tomorrow will be handled sooner than a task for 42 workers starting in 2 months.*

---

# Duration

Depending on your levels in Golang, databases or back-end architecture, the implementation can be more or less time consuming.

We consider that having a functional **US1.0**, **US1.1**, **US2.0** and **US3.0** is the **bare minimum** and we think that 48 hours (2 days) should be a maximum. If you didn’t have the time to do everything you wanted, just write us what you had in mind, it’ll be sufficient!

# Mindset

The purpose of this test is not to track how fast you can go but more to understand how you understand/solve problems through the usage of your experience(s) and learning(s) as well as technology.

Please, make sure that the test is easily launchable. The easier it’ll be to review your test, the faster we’ll provide you with an answer.

# Bonus

1. Dockerize your project
2. Think about scalability. How would you handle a list of 1000 tasks? Or 1000 tasks with 1000 slots each?
3. Write your own tech specs to explain your technical choices
4. Authentication system
5. Feel free to add any feature that could bring more value to your project! 

---

# Test

For the purpose of this test, you’ll be working on a tool that is firstly useful for step 6 (it can be useful for other steps as well, such as steps 2. - 3. - 4. - 5.).

The goal of this tool is to display tasks in a prioritized list, clearly highlighting key information to help our operational team to handle them as fast as possible.

You’ll need to provide the right data to display information about

- The **tasks** - entities holding data about the need itself
- The **shifts** - entities holding the different dates at which the tasks start and end
- The **slots** - individual entities linked to a shift. They contain information about the sider (temp worker) working on the shift

---

# Functional specifications

## [EPIC1] Display

- [ ]  **US 1.0 - As a user, I want a list of the tasks to be displayed**

Including:

- The task name
- The organisation name
- The organisation picture
- The number of slots, filled and available
- The number of applicants

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/3e2c3b34-cb7b-4ea8-a068-81a5be569058/Untitled.png)

- [ ]  **US 1.1 - As a user, I want to be able to filter the tasks depending on their status**

Status can be one of:

- Upcoming
- Ongoing
- Done
- [ ]  **US 1.2 (BONUS) - As a user, when I hover over the slots part, I want to be able to get information about the different shifts of each task**

A shift must contain the following information:

- The start date
- The end date
- The number of slots available for this shift
- The number of slots filled for this shift
    
    ![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/c6fa8cbb-0824-4a3c-8db2-b810796bdcf6/Untitled.png)
    
- [ ]  **US 1.3 (BONUS) - As a user, when I hover over the company name, I want the location of the task to be displayed**

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/32d44918-627c-4695-a153-e2c726adae57/Untitled.png)

- [ ]  **US 1.4 (BONUS) - As a user, I want the Ops member assigned to this task to be displayed**

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/47d431e7-619c-41db-8141-78ceae4641b4/Untitled.png)

## [EPIC2] Edit

- [ ]  **US 2.0 - As a user, I want to be able to modify the Ops member assigned to this task**

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/8c49decb-aee8-460d-a7d0-dbff619e91ac/Untitled.png)

# [EPIC3] Open question

- [ ]  **US 3.0 - Tasks are not available to the siders right away after they’re created. Which functional system would you consider to make them available through the mobile application? You can send us your written answer/schema alongside your project.**

# Technical specifications

## What you must know

For the purpose of this test, we will provide you:

- The data required to build your feature
- The payload expected by the front-end

In the meantime, you’ll be in charge of taking decision about:

- The type of database you want to use (sql, nosql)
- The API protocol

Here are some specifications:

- The project must be done in Go
- You can use any (stable) version of Go you want
- You can choose any database type you want
- You have to make a micro-services architecture
- You must provide a simple way to build and run it

## Specifications

### Endpoint between the front-end and your back-end service:

`GET /tasks` *Note: this is the base payload, for some US you’ll need to add some fields or parameters.*

```jsx
{
    "tasks": [
        {
            "id": "ta_FM3fQr4QtHHE",
            "name": "Agent logistique en magasinage (F/H) - Affréteur International",
            "organisation": {
                "name": "Recargo",
                "address": "70115 Rue Eugène Pottier , Noisy-le-Sec, Tarn-et-Garonne (82) 14426",
                "pictureUrl": "https://picsum.photos/400/400"
            },
            "shifts": [
                {
                    "id": "sh_2HwbZYgZAOefRp2I6l0J4DwpT1g",
                    "startDate": "2023-01-21T10:00:00.000Z",
                    "endDate": "2023-01-21T18:00:00.000Z",
                    "slots": {
                        "filled": 1,
                        "total": 1
                    },
                }
            ]
        }
    ]
}
```

`PATCH /tasks`

```jsx
{
	    "assigneeId": "3fb4dbcf1a4f2b0316d3da5c9343f12c1a55f524ff73506561eb043823c9e088"
}
```

### Provided data sets

In the archive you will find 5 different data set (or collections). Here is a short description of each of them:

- `tasks` contains all the informations for a task : general informations on the task, task name, temp workers who have applied for the task, etc
- `shifts` contains information about the different dates available for the tasks
- `slots` contains information about who is working on a task
- `users` contains the general informations about users, their name, email, etc
- `organisations` contains informations about the orgas like their address or their picture url

---

# Resources

[dump.zip](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/9804a0cc-40e4-44a0-b6a4-e64392a45582/dump.zip)
