# MailUpdater

The mailupdater takes two email addresses (old and new), checks
the database for the existence of users with either of the emails.

If both emails addresses are found, return and error message and 
halt further processing or updates.

If the new email address is not found in the database, then update
the user email to the new, essentially replacing the old one.

Lastly we notify the support team of the changes and update a spreadsheet
with the updated information.