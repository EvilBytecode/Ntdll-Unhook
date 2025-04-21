## Ntdll Unhook (DLL Unhooking)

Endpoint Detection and Response (EDR) tools use a technique known as hooking to monitor sensitive system functions within the DLLs of loaded processes. Hooking is a method of live-patching system DLLs, enabling EDRs to intercept the flow of a program and evaluate its legitimacy.

Here's how it works: EDRs modify the first instructions of the functions within the DLLs. When these functions are called, the program's execution flow is diverted to the EDR's code (housed within a DLL loaded by the EDR in the program). In this redirected state, the EDR can inspect the function's arguments to determine whether their usage is legitimate or potentially malicious. If the usage is deemed legitimate, the EDR restores the program's execution flow, allowing the function to proceed as normal.

However, to evade detection by an EDR, malware can employ a method known as "unhooking." This process involves restoring the entire DLL code section (.text) to its original state. To accomplish this, malware needs access to an unmodified (unhooked) DLL, which it can acquire in several ways:

1 - directly from the system, which can potentially be detected via an open handle;

2 - by opening a remote file, which requires the malware author to host a DLL matching the OS version of the target system remotely;

3 - by initiating a suspended process and retrieving the content of its DLL before it gets hooked.

Typically, the DLL most commonly hooked/unhooked is NTDLL.dll, as it is the closest to the kernel. However, some EDRs may also hook APIs contained in higher-level DLLs, such as kernel32.dll or user32.dll.

## setup:
- Run build.bat, it will do everything for you.


## License
This project is licensed under the MIT License. See the LICENSE file for details.