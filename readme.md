

# Go Producer-Consumer and String Reversal Example

This Go project demonstrates a combination of two common programming concepts: 
1. A **Producer-Consumer** pattern with rate-limited consumption.
2. **String Reversal** logic for a list of strings.

## Features
1. **Producer-Consumer Pattern**: 
   - A producer generates items at a faster rate and sends them to a channel.
   - A consumer retrieves items from the channel at a controlled, rate-limited pace.

2. **String Reversal**: 
   - Strings are reversed from an input list unless the string length is exactly 6 characters, in which case the operation stops.

## Prerequisites

- Go 1.16 or higher installed. You can download it from [here](https://golang.org/dl/).

## How to Run the Project

### 1. Clone the Repository

You can clone this repository using the following command:

```bash
gh repo clone bsagute/all-files-data
cd all-files-data
```

Alternatively, you can clone it using HTTPS:

```bash
git clone https://github.com/bsagute/all-files-data.git
cd all-files-data
```

### 2. Run the Go Program

After navigating to the directory containing the `main.go` file, you can run the Go program using the following command:

```bash
go run main.go
```

### Expected Output

The program will output the following:

1. It will **produce** and **consume** items:
    ```plaintext
    Produced item: 1
    Consumed item: 1
    Produced item: 2
    Consumed item: 2
    ...
    ```

2. It will reverse the strings from the list:
    ```plaintext
    Reversed: olleH
    Reversed: CBA
    Reversed: ZYX
    Reversed: MKK
    ```

3. Finally, a message will be displayed when all operations are completed:
    ```plaintext
    All items produced, consumed, and string reversal completed.
    ```

## Code Breakdown

- **Producer-Consumer Logic**: 
    - The `producer` function simulates a burst of production by sending items to a channel.
    - The `consumer` function processes the items from the channel at a rate-limited pace (2 items per second).

- **String Reversal Logic**:
    - The `reverseStrings` function takes a list of strings and reverses each one using the `DoRev` function. It stops if a string has a length of exactly 6 characters.

## File Structure

```plaintext
├── main.go
└── README.md
```

## Customization

- **Modify total items**: You can change the number of items produced by modifying the `totalItems` constant in the `main.go` file.
- **Change rate limit**: The rate of consumption (in items per second) can be adjusted by changing the `rateLimit` constant.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

### Enjoy coding with Go! If you have any issues, feel free to open a new issue in the repository.

---
