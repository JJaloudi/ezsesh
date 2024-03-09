# EZSesh - The Simple Session Store


I created EZSesh to handle internal application session management. When using other FOSS session management tools, I found their implementation to hide away too much abstraction. 
The development of this package is to avoid all of that. Writing your own stores should be easy, you should be able to dictate everything down to the hashing and storing of cookies.

I want to ultimately leave every step of cookie generation up to the developer, while maintaining Go's concurrency & following industry standard session management practices.

----

# FYI

Currently, this package is a WIP, primarily the development now is focused on handling my internal usecases, however, the direction during development is to create tons of genericism so that DX can be 
tailored to any project and developer. But right now, the lib does not include any documentation in the code. This will change over the coming weeks for anyone that manages to find this repo :)


As of now, I recommend against using this package, as there are multiple cybersec items that need to be tested and cleared, like session fixation, hijacking, timeouts & other basic stuff. As of now,
this is used to maintain progress on this repository.

Current plans for development:
- [x] Create SQLX oriented store
- [x] Utilize fixed time hashing methods
- [ ] Generic methods for every step of cookie generation
- [ ] Write base stores with an emphasis on concurrency

This list will continue to grow with both finished and unfinished future tasks.
