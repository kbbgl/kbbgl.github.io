# Linux Kernel Development

Summarized information from [Linux Kernel Development, 3rd Edition](https://www.oreilly.com/library/view/linux-kernel-development/9780768696974/).

The kernel is the innermost portion of the operating system.

It provides basic services for all other parts of the system, manages hardware and distributes system resources.

Typical components of a kernel are interrupt handlers to service interrupt requests, a scheduler to share processor time, a memory management system to manage process address spaces and system services such as networking and interprocess communication.

Applications running on the system communicate with the kernel via _system calls_. The system call interface instructs the kernel to carry out tasks on behalf of the application. The instructions are running in _kernel-space_.

The kernel manages the system's hardware. When hardware wants to communicate with the system, it issues an numbered interrupt to the process which interrupts the kernel. The interrupt number controls which interrupt handler should process the request. Interrupts are run in an _interrupt context_.

Each processor is doing one of three things at any given moment:

- In user-space, executing user code in a process.
- In kernel-space, in process context, executing on behalf of a specific process.
- In kernel-space, in interrupt context, not associated with a process, handling an interrupt.

![contexts](https://www.form3.tech/_prismic-media/6ad9db5a55ad1b8ab8616158bb1bc16f6cb34171147534f791fed46cf2d3b6c7.png)
