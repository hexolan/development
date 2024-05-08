using Microsoft.AspNetCore.Identity;

namespace Formulator.Infrastructure.Identity
{
    public class ApplicationUser : IdentityUser
    {
        public string? DisplayName { get; set; }
    }
}
