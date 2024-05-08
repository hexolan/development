using Formulator.Core.Entities;

namespace Formulator.Infrastructure.Data
{
    internal class ApplicationDbInitialiser
    {
        public static void Initialise(ApplicationDbContext context)
        {
            // todo: instead ensure migrations
            context.Database.EnsureCreated();

            /*
            // Check if the database has already been seeded
            if (context.Users.Any())
            {
                return;
            }
            */
        }
    }
}
