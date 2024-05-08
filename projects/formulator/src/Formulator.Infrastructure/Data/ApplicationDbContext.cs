using Formulator.Core.Entities;
using Formulator.Infrastructure.Identity;
using Microsoft.AspNetCore.Identity.EntityFrameworkCore;
using Microsoft.EntityFrameworkCore;

namespace Formulator.Infrastructure.Data
{
    public class ApplicationDbContext(DbContextOptions<ApplicationDbContext> options) : IdentityDbContext<ApplicationUser>(options)
    {
        public DbSet<Form> Forms { get; set; }
        public DbSet<FormQuestion> FormQuestions { get; set; }

        public DbSet<FormSubmission> FormSubmissions { get; set; }
        public DbSet<FormSubmissionResponse> FormSubmissionResponses { get; set; }

        protected override void OnModelCreating(ModelBuilder modelBuilder)
        {
            modelBuilder.Ignore<User>();
            modelBuilder.Entity<FormSubmissionResponse>()
                .HasKey(m => new { m.SubmissionId, m.QuestionId });

            base.OnModelCreating(modelBuilder);
        }
    }
}
